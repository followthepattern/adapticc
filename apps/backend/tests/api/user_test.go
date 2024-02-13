package test_api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/api/middlewares"
	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/features/user"
	"github.com/followthepattern/adapticc/mocks"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/tests/datagenerator"
	"github.com/followthepattern/adapticc/tests/sqlexpectations"
	"github.com/followthepattern/adapticc/types"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"

	"github.com/DATA-DOG/go-sqlmock"
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

var _ = Describe("User graphql queries", func() {
	var (
		mdb         *sql.DB
		mock        sqlmock.Sqlmock
		cfg         config.Config
		jwtKeys     config.JwtKeyPair
		handler     http.Handler
		mockCtrl    *gomock.Controller
		mockCerbos  *mocks.MockClient
		contextUser auth.AuthUser
		roleIDs     []string
	)

	BeforeEach(func() {
		cfg = config.Config{
			Server: config.Server{
				HmacSecret:            "test",
				GraphqlSchemaFilepath: "./../../api/graphql/schema/schema.graph",
			}}
		var err error
		mdb, mock, err = sqlmock.New()
		Expect(err).To(BeNil())

		jwtKeys, err = getMockJWTKeys()
		Expect(err).ShouldNot(HaveOccurred())

		mockCtrl = gomock.NewController(GinkgoT())
		mockCerbos = mocks.NewMockClient(mockCtrl)

		ac := accesscontrol.Config{
			Kind:   "user",
			Cerbos: mockCerbos,
		}.Build()

		handler = NewMockHandler(ac, nil, mdb, cfg, jwtKeys)

		password := datagenerator.String(18)
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		Expect(err).Should(BeNil())
		contextUser = datagenerator.NewRandomAuthUser(passwordHash)
		roleIDs = []string{datagenerator.String(18), datagenerator.String(18)}
	})

	Context("Single", func() {
		It("Success", func() {
			user := datagenerator.NewRandomUser()

			queryTemplate := `
			query {
				users {
					single(id: "%v") {
						id
						email
					}
				}
			}`
			query := fmt.Sprintf(queryTemplate, user.ID)
			graphRequest := graphqlRequest{
				Query: query,
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectGetUserByID(mock, user, user.ID)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))
			Expect(testResponse.Data.Users.Single.ID.Data).To(Equal(user.ID.Data))
		})
	})

	Context("List", func() {
		It("Success default query", func() {
			users := []user.UserModel{datagenerator.NewRandomUser(), datagenerator.NewRandomUser(), datagenerator.NewRandomUser()}

			listRequestParams := user.UserListRequestParams{
				Pagination: models.Pagination{
					PageSize: types.UintFrom(5),
					Page:     types.UintFrom(1),
				},
				Filter: models.ListFilter{
					Search: types.StringFrom(datagenerator.RandomEmail(8, 8)),
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

			query := fmt.Sprintf(queryTemplate, listRequestParams.Filter.Search)
			graphRequest := graphqlRequest{
				Query: query,
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectUsers(mock, users, listRequestParams)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(BeEmpty())
			Expect(testResponse.Data.Users.List.Data).To(HaveLen(len(users)))
			for i := range testResponse.Data.Users.List.Data {
				Expect(testResponse.Data.Users.List.Data[i].ID.Data).To(Equal(users[i].ID.Data))
				Expect(testResponse.Data.Users.List.Data[i].Email.Data).To(Equal(users[i].Email.Data))
				Expect(testResponse.Data.Users.List.Data[i].FirstName.Data).To(Equal(users[i].FirstName.Data))
				Expect(testResponse.Data.Users.List.Data[i].LastName.Data).To(Equal(users[i].LastName.Data))
			}

			Expect(mock.ExpectationsWereMet()).To(BeNil())
		})

		It("Success withouth page and pageSize params", func() {
			users := []user.UserModel{datagenerator.NewRandomUser(), datagenerator.NewRandomUser(), datagenerator.NewRandomUser()}

			listRequestParams := user.UserListRequestParams{
				Filter: models.ListFilter{
					Search: types.StringFrom(datagenerator.RandomEmail(8, 8)),
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
			query := fmt.Sprintf(queryTemplate, listRequestParams.Filter.Search)
			graphRequest := graphqlRequest{
				Query: query,
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectUsersWithoutPaging(mock, users, listRequestParams)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(BeEmpty())
			Expect(testResponse.Data.Users.List.Data).To(HaveLen(len(users)))
			for i := range testResponse.Data.Users.List.Data {
				Expect(testResponse.Data.Users.List.Data[i].ID.Data).To(Equal(users[i].ID.Data))
				Expect(testResponse.Data.Users.List.Data[i].Email.Data).To(Equal(users[i].Email.Data))
				Expect(testResponse.Data.Users.List.Data[i].FirstName.Data).To(Equal(users[i].FirstName.Data))
				Expect(testResponse.Data.Users.List.Data[i].LastName.Data).To(Equal(users[i].LastName.Data))
			}

			Expect(mock.ExpectationsWereMet()).ShouldNot(HaveOccurred())
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

			graphRequest := graphqlRequest{
				Query: query,
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectGetUserByID(mock, user, contextUser.ID)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Profile.ID.Data).To(Equal(user.ID.Data))
		})
	})

	Context("Create", func() {
		It("Succeeds", func() {
			user := datagenerator.NewRandomUser()

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

			query := fmt.Sprintf(queryTemplate, user.Email, user.FirstName, user.LastName)
			graphRequest := graphqlRequest{
				Query: query,
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.CREATE).Return(true, nil)
			sqlexpectations.ExpectCreateUser(mock, contextUser.ID, user)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Create.Code).To(Equal(int32(http.StatusCreated)))
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
			user := datagenerator.NewRandomUser()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(graphql, user.ID, user.FirstName, user.LastName),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.UPDATE).Return(true, nil)
			sqlexpectations.ExpectUpdateUser(mock, contextUser.ID, user)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Update.Code).To(Equal(int32(http.StatusOK)))
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
			user := datagenerator.NewRandomUser()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(graphql, user.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.DELETE).Return(true, nil)
			sqlexpectations.ExpectDeleteUser(mock, user.ID)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Delete.Code).To(Equal(int32(http.StatusOK)))
		})
	})
})
