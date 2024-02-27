package test_integration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/graph-gophers/graphql-go/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type graphqlAuthResponse struct {
	Data   authenticationData `json:"data"`
	Errors []*errors.QueryError
}

type authenticationData struct {
	Authentication authentication `json:"authentication"`
}

type authentication struct {
	Register auth.RegisterResponse `json:"register"`
	Login    auth.LoginResponse    `json:"login"`
}

var _ = Describe("Auth queries", Ordered, func() {
	var (
		backendAbsolutePath, _ = filepath.Abs("./../../")

		testDir *dagger.Directory

		ctx     context.Context
		client  *dagger.Client
		backend *dagger.Service
	)

	BeforeAll(func() {
		var err error
		ctx = context.Background()

		client, err = dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		Expect(err).To(BeNil())

		backendDirectory := client.Host().Directory(backendAbsolutePath)

		builder := client.Container().From("golang:latest")
		builder = builder.WithDirectory("/backend", backendDirectory).WithWorkdir("/backend")

		outputPath := "out/"
		builder = builder.WithExec([]string{"go", "build", "-o", outputPath, "./cmd/adapticc"})

		output := builder.Directory(outputPath)
		_, err = output.Export(ctx, filepath.Join(backendAbsolutePath, outputPath))
		Expect(err).To(BeNil())

		testDataDir := client.Host().Directory("./testdata")

		database := client.Container().From("postgres:latest").
			WithEnvVariable("POSTGRES_DB", "adapticc").
			WithDirectory("/docker-entrypoint-initdb.d", testDataDir).
			WithEnvVariable("POSTGRES_USER", "adapticcuser").
			WithEnvVariable("POSTGRES_PASSWORD", "dbpass").
			WithExec([]string{"postgres"}).
			WithExposedPort(5432).
			AsService()

		backend = client.Container().From("golang:1.21").
			WithServiceBinding("adapticc_db", database).
			WithDirectory("/backend", backendDirectory).
			WithWorkdir("/backend").
			WithExec([]string{"./out/adapticc"}).
			WithExposedPort(8080).
			AsService()

	})

	Context("Login", func() {
		BeforeEach(func() {
			testDir = client.Host().Directory(".")
		})

		It("succeeds to login with admin", func() {
			var (
				queryTemplate = `
							mutation {
								authentication {
									login(email: "%v", password: "%v") {
										jwt
										expires_at
									}
								}
							}`

				loginEmail = "admin@admin.com"
				password   = "Admin@123!"
			)

			query := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, loginEmail, password),
			}

			requestBody, _ := json.Marshal(query)

			out, err := client.Container().From("golang:1.21").
				WithServiceBinding("backend", backend).
				WithDirectory("/httpClient", testDir).
				WithWorkdir("/httpClient").
				WithExec([]string{"go", "run", "./http_tester/client.go", http.MethodPost, graphQLURL, string(requestBody)}).
				Stdout(ctx)

			Expect(err).Should(BeNil())

			response := graphqlAuthResponse{}

			err = json.Unmarshal([]byte(out), &response)
			Expect(err).Should(BeNil())

			Expect(response.Errors).Should(BeEmpty())

			Expect(response.Data.Authentication.Login.JWT).ShouldNot(BeEmpty())
		})
	})

	AfterAll(func() {
		client.Close()
	})

})
