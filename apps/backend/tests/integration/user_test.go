package test_integration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"dagger.io/dagger"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("User queries", Ordered, func() {
	var (
		backendAbsolutePath, _ = filepath.Abs("./../../")
		testDir                *dagger.Directory

		ctx     context.Context
		client  *dagger.Client
		backend *dagger.Service

		graphQLURL = "http://backend:8080/graphql"
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

		database := client.Container().From("postgres:latest").
			WithEnvVariable("POSTGRES_DB", "adapticc").
			WithEnvVariable("POSTGRES_USER", "adapticcuser").
			WithEnvVariable("POSTGRES_PASSWORD", "dbpass").
			WithExec([]string{"postgres"}).
			WithExposedPort(5432).
			AsService()

		backend = client.Container().From("golang:1.21").
			WithServiceBinding("db", database).
			WithDirectory("/backend", backendDirectory).
			WithWorkdir("/backend").
			WithExec([]string{"./out/adapticc"}).
			WithExposedPort(8080).
			AsService()

	})

	Context("Single", func() {
		BeforeEach(func() {
			testDir = client.Host().Directory(".")
		})

		FIt("succeeds to return with a user", func() {
			queryStr := `
			query {
				users {
					profile {
						id
						email
						firstName
						lastName
					}
				}
			}`

			query := graphqlRequest{
				Query: queryStr,
			}

			requestBody, _ := json.Marshal(query)

			out, err := client.Container().From("golang:1.21").
				WithServiceBinding("backend", backend).
				WithDirectory("/httpClient", testDir).
				WithWorkdir("/httpClient").
				WithExec([]string{"go", "run", "./http_tester/client.go", http.MethodPost, graphQLURL, string(requestBody)}).
				Stdout(ctx)

			Expect(err).Should(BeNil())

			fmt.Println(out)
		})
	})

	AfterAll(func() {
		client.Close()
	})

})
