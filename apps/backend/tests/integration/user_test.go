package test_integration

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"dagger.io/dagger"
	"github.com/followthepattern/adapticc/features/user"
	"github.com/followthepattern/adapticc/models"
	"github.com/graph-gophers/graphql-go/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type graphqlUserResponse struct {
	Data   userData             `json:"data"`
	Errors []*errors.QueryError `json:"errors,omitempty"`
}

type userData struct {
	Users users `json:"users"`
}

type users struct {
	Single  user.UserModel                      `json:"single,omitempty"`
	Profile user.UserModel                      `json:"profile,omitempty"`
	List    models.ListResponse[user.UserModel] `json:"list,omitempty"`
	Create  models.ResponseStatus               `json:"create,omitempty"`
	Update  models.ResponseStatus               `json:"update,omitempty"`
	Delete  models.ResponseStatus               `json:"delete,omitempty"`
}

var _ = Describe("User queries", Ordered, func() {
	var (
		backendAbsolutePath, _ = filepath.Abs("./../../")

		testDir *dagger.Directory

		ctx     context.Context
		client  *dagger.Client
		backend *dagger.Service

		jwtToken     string
		userResponse graphqlUserResponse

		testUserEmail    = "admin@admin.com"
		testUserPassword = "Admin@123!"
	)

	AssertSucceedsLogin := func(email, password string) {
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
		)

		query := graphqlRequest{
			Query: fmt.Sprintf(queryTemplate, email, password),
		}

		requestBody, _ := json.Marshal(query)

		out, err := client.Container().From(GolangImage).
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

		jwtToken = response.Data.Authentication.Login.JWT
	}

	BeforeAll(func() {
		var err error
		ctx = context.Background()

		client, err = dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
		Expect(err).To(BeNil())

		backendDirectory := client.Host().Directory(backendAbsolutePath)
		testDir = client.Host().Directory(".")

		builder := client.Container().From(GolangImage)
		builder = builder.WithDirectory("/backend", backendDirectory).WithWorkdir("/backend")

		outputPath := "out/"
		builder = builder.WithExec([]string{"go", "build", "-o", outputPath, "./cmd/adapticc"})

		output := builder.Directory(outputPath)
		_, err = output.Export(ctx, filepath.Join(backendAbsolutePath, outputPath))
		Expect(err).To(BeNil())

		testDataDir := client.Host().Directory("./testdata")

		database := client.Container().From(PostgresImage).
			WithEnvVariable("POSTGRES_DB", "adapticc").
			WithDirectory("/docker-entrypoint-initdb.d", testDataDir).
			WithEnvVariable("POSTGRES_USER", "adapticcuser").
			WithEnvVariable("POSTGRES_PASSWORD", "dbpass").
			WithExec([]string{"postgres"}).
			WithExposedPort(5432).
			AsService()

		backend = client.Container().From(GolangImage).
			WithServiceBinding("adapticc_db", database).
			WithDirectory("/backend", backendDirectory).
			WithWorkdir("/backend").
			WithExec([]string{"./out/adapticc"}).
			WithExposedPort(8080).
			AsService()

		AssertSucceedsLogin(testUserEmail, testUserPassword)

	})

	Context("Profile", func() {
		It("returns with the signed in user", func() {
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

			out, err := client.Container().From(GolangImage).
				WithServiceBinding("backend", backend).
				WithDirectory("/httpClient", testDir).
				WithWorkdir("/httpClient").
				WithExec([]string{"go", "run", "./http_tester/client.go", http.MethodPost, graphQLURL, string(requestBody), jwtToken}).
				Stdout(ctx)

			Expect(err).Should(BeNil())
			json.Unmarshal([]byte(out), &userResponse)

			Expect(userResponse.Data.Users.Profile.ID.Data).ShouldNot(BeEmpty())
			Expect(userResponse.Data.Users.Profile.Email.Data).To(Equal(testUserEmail))
		})
	})

	AfterAll(func() {
		client.Close()
	})

})
