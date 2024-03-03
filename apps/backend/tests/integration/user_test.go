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
	"github.com/followthepattern/adapticc/tests/datagenerator"
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

		cerbosDir := client.Host().Directory(filepath.Join(backendAbsolutePath, "cerbos"))

		cerbos := client.Container().From(CerbosImage).
			WithDirectory("/data", cerbosDir).
			WithExec([]string{"server", "--config=/data/.cerbos.yaml"}).
			WithExposedPort(3592).
			AsService()

		backend = client.Container().From(GolangImage).
			WithServiceBinding("adapticc_db", database).
			WithServiceBinding("cerbos", cerbos).
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

	Context("Single", func() {
		It("returns with a user by id", func() {
			userID := "613254df-c779-479c-9d76-b8036e342979"

			queryTemplate := `
				query {
					users {
						single(id: "%v") {
							id
							email
						}
					}
				}`

			query := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, userID),
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

			Expect(userResponse.Data.Users.Single.ID.Data).Should(Equal(userID))
			Expect(userResponse.Data.Users.Single.Email.Data).To(Equal(testUserEmail))
		})
	})

	Context("List", func() {
		It("returns with users", func() {
			queryStr := `
				query {
					users {
						list (
							filter: { search: "" }
						) {
							page
							pageSize
							count
							data {
								id
								email
								firstName
								lastName
							}
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

			Expect(userResponse.Data.Users.List.Count.Data).ShouldNot(Equal(0))
			Expect(userResponse.Data.Users.List.Data).Should(HaveEach(HaveField("ID.Data", Not(BeEmpty()))))
			Expect(userResponse.Data.Users.List.Data).Should(HaveEach(HaveField("Email.Data", Not(BeEmpty()))))
			Expect(userResponse.Data.Users.List.Data).Should(HaveEach(HaveField("FirstName.Data", Not(BeEmpty()))))
			Expect(userResponse.Data.Users.List.Data).Should(HaveEach(HaveField("LastName.Data", Not(BeEmpty()))))
		})
	})

	Context("Create", func() {
		It("creates a new user", func() {
			createdUser := datagenerator.NewRandomUser()

			queryTemplate := `
					mutation {
						users {
							create (model: {
								email: "%s"
								firstName: "%s"
								lastName: "%s"
							}) {
								code
							}
						}
					}`

			query := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, createdUser.Email.Data, createdUser.FirstName.Data, createdUser.LastName.Data),
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

			Expect(userResponse.Data.Users.Create.Code).Should(Equal(int32(http.StatusCreated)))
		})
	})

	AfterAll(func() {
		client.Close()
	})

})
