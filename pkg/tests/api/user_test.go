package test_api

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/api/middlewares"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/tests/datagenerator"
	"github.com/followthepattern/adapticc/pkg/tests/sqlexpectations"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"github.com/golang-jwt/jwt/v4"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/graphql-go/errors"

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
	Single  models.User                      `json:"single,omitempty"`
	Profile models.User                      `json:"profile,omitempty"`
	List    models.ListResponse[models.User] `json:"list,omitempty"`
	Update  resolvers.ResponseStatus         `json:"update,omitempty"`
	Delete  resolvers.ResponseStatus         `json:"delete,omitempty"`
}

var _ = Describe("User graphql queries", func() {
	var (
		mdb     *sql.DB
		mock    sqlmock.Sqlmock
		ctx     context.Context
		cfg     config.Config
		handler http.Handler
	)

	BeforeEach(func() {
		ctx = context.Background()
		cfg = config.Config{
			Server: config.Server{
				HmacSecret: "test",
			}}
		var err error
		mdb, mock, err = sqlmock.New()
		Expect(err).To(BeNil())

		handler = NewMockHandler(ctx, mdb, cfg)

	})

	Context("Single", func() {
		It("Success", func() {
			queryTemplate := `
			query {
				users {
					single(id: "%v") {
						id
						email
					}
				}
			}`

			user := datagenerator.NewRandomUser()

			query := fmt.Sprintf(queryTemplate, *user.ID)

			graphRequest := graphqlRequest{
				Query: query,
			}

			request, _ := json.Marshal(graphRequest)

			sqlexpectations.ExpectGetUserByID(mock, user, *user.ID)

			testResponse := &graphqlUserResponse{}

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))
			Expect(*testResponse.Data.Users.Single.ID).To(Equal(*user.ID))
		})
	})

	Context("List", func() {
		It("Success default query", func() {
			listRequestParams := models.UserListRequestParams{
				Pagination: models.Pagination{
					PageSize: pointers.ToPtr[uint](5),
					Page:     pointers.ToPtr[uint](1),
				},
				Filter: models.ListFilter{
					Search: pointers.ToPtr("email@email.com"),
				},
			}

			queryTemplate := `
			query {
				users {
					list (
						pagination: { pageSize: 5, page: 1 }
						filter: { search: "%s" }
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

			query := fmt.Sprintf(queryTemplate, *listRequestParams.Filter.Search)

			graphRequest := graphqlRequest{
				Query: query,
			}

			request, _ := json.Marshal(graphRequest)

			users := []models.User{datagenerator.NewRandomUser(), datagenerator.NewRandomUser(), datagenerator.NewRandomUser()}

			sqlexpectations.ExpectUsers(mock, "", users, listRequestParams)

			testResponse := &graphqlUserResponse{}

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(BeEmpty())
			Expect(testResponse.Data.Users.List.Data).To(HaveLen(len(users)))
			for i := range testResponse.Data.Users.List.Data {
				Expect(*testResponse.Data.Users.List.Data[i].ID).To(Equal(*users[i].ID))
				Expect(*testResponse.Data.Users.List.Data[i].Email).To(Equal(*users[i].Email))
				Expect(*testResponse.Data.Users.List.Data[i].FirstName).To(Equal(*users[i].FirstName))
				Expect(*testResponse.Data.Users.List.Data[i].LastName).To(Equal(*users[i].LastName))
			}

			Expect(mock.ExpectationsWereMet()).To(BeNil())
		})

		It("Success withouth page and pageSize params", func() {
			listRequestParams := models.UserListRequestParams{
				Filter: models.ListFilter{
					Search: pointers.ToPtr("email@email.com"),
				},
			}

			queryTemplate := `
			query {
				users {
					list (
						filter: { search: "%s" }
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

			query := fmt.Sprintf(queryTemplate, *listRequestParams.Filter.Search)

			graphRequest := graphqlRequest{
				Query: query,
			}

			request, _ := json.Marshal(graphRequest)

			users := []models.User{datagenerator.NewRandomUser(), datagenerator.NewRandomUser(), datagenerator.NewRandomUser()}

			sqlexpectations.ExpectUsersWithoutPaging(mock, "", users, listRequestParams)

			testResponse := &graphqlUserResponse{}

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(BeEmpty())
			Expect(testResponse.Data.Users.List.Data).To(HaveLen(len(users)))
			for i := range testResponse.Data.Users.List.Data {
				Expect(*testResponse.Data.Users.List.Data[i].ID).To(Equal(*users[i].ID))
				Expect(*testResponse.Data.Users.List.Data[i].Email).To(Equal(*users[i].Email))
				Expect(*testResponse.Data.Users.List.Data[i].FirstName).To(Equal(*users[i].FirstName))
				Expect(*testResponse.Data.Users.List.Data[i].LastName).To(Equal(*users[i].LastName))
			}

			Expect(mock.ExpectationsWereMet()).To(BeNil())
		})
	})

	Context("Profile", func() {
		var query string = `
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

		It("Success", func() {
			user := datagenerator.NewRandomUser()

			expiresAt := time.Now().Add(time.Hour * 24)
			token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"ID":        *user.ID,
				"email":     *user.Email,
				"firstName": *user.FirstName,
				"lastName":  *user.LastName,
				"expiresAt": expiresAt,
			})

			tokenString, err := token.SignedString([]byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			graphRequest := graphqlRequest{
				Query: query,
			}

			request, _ := json.Marshal(graphRequest)

			sqlexpectations.ExpectGetUserByID(mock, user, *user.ID)

			testResponse := &graphqlUserResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(*testResponse.Data.Users.Profile.ID).To(Equal(*user.ID))
		})
	})

	Context("Update", func() {
		var graphql string = `mutation {
			users {
				update (id: "%v", model: {
					firstName: "%s"
					lastName: "%s"
				}) {
					code
				}
			}
		}`

		It("Success", func() {
			contextUser := datagenerator.NewRandomUser()
			user := datagenerator.NewRandomUser()

			expiresAt := time.Now().Add(time.Hour * 24)
			token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"ID":        *contextUser.ID,
				"email":     *contextUser.Email,
				"firstName": *contextUser.FirstName,
				"lastName":  *contextUser.LastName,
				"expiresAt": expiresAt,
			})

			tokenString, err := token.SignedString([]byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(graphql, *user.ID, *user.FirstName, *user.LastName),
			}

			request, _ := json.Marshal(graphRequest)

			sqlexpectations.ExpectUpdateUser(mock, *contextUser.ID, user)

			testResponse := &graphqlUserResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Update.Code).To(Equal(resolvers.NewUint(200)))
		})
	})

	Context("Delete", func() {
		var graphql string = `mutation {
			users {
				delete (id: "%v") {
					code
				}
			}
		}`

		It("Success", func() {
			contextUser := datagenerator.NewRandomUser()
			user := datagenerator.NewRandomUser()

			expiresAt := time.Now().Add(time.Hour * 24)
			token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
				"ID":        *contextUser.ID,
				"email":     *contextUser.Email,
				"firstName": *contextUser.FirstName,
				"lastName":  *contextUser.LastName,
				"expiresAt": expiresAt,
			})

			tokenString, err := token.SignedString([]byte(cfg.Server.HmacSecret))
			if err != nil {
				panic(err)
			}

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(graphql, *user.ID),
			}

			request, _ := json.Marshal(graphRequest)

			sqlexpectations.ExpectDeleteUser(mock, "", user)

			testResponse := &graphqlUserResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, tokenString)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Delete.Code).To(Equal(resolvers.NewUint(200)))
		})
	})
})
