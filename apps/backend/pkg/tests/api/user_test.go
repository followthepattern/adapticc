package test_api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/followthepattern/adapticc/mocks"
	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/api/middlewares"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/services"
	"github.com/followthepattern/adapticc/pkg/tests/datagenerator"
	"github.com/followthepattern/adapticc/pkg/tests/sqlexpectations"
	"github.com/followthepattern/adapticc/pkg/utils/pointers"
	"github.com/golang/mock/gomock"

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
	Create  resolvers.ResponseStatus         `json:"create,omitempty"`
	Update  resolvers.ResponseStatus         `json:"update,omitempty"`
	Delete  resolvers.ResponseStatus         `json:"delete,omitempty"`
}

var _ = Describe("User graphql queries", func() {
	var (
		mdb        *sql.DB
		mock       sqlmock.Sqlmock
		cfg        config.Config
		handler    http.Handler
		mockCtrl   *gomock.Controller
		mockCerbos *mocks.MockClient
	)

	BeforeEach(func() {
		cfg = config.Config{
			Server: config.Server{
				HmacSecret: "test",
			}}
		var err error
		mdb, mock, err = sqlmock.New()
		Expect(err).To(BeNil())

		mockCtrl = gomock.NewController(GinkgoT())
		mockCerbos = mocks.NewMockClient(mockCtrl)

		ac := accesscontrol.Config{
			Kind:   "user",
			Cerbos: mockCerbos,
		}.Build()

		handler = NewMockHandler(ac, nil, mdb, cfg)

	})

	Context("Single", func() {
		It("Success", func() {
			user := datagenerator.NewRandomUser()
			role1 := datagenerator.NewRandomRole()

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

			tokenString, err := services.GenerateTokenStringFromUser(user, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, user.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectGetUserByID(mock, user, user.ID)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))
			Expect(testResponse.Data.Users.Single.ID).To(Equal(user.ID))
		})
	})

	Context("List", func() {
		It("Success default query", func() {
			users := []models.User{datagenerator.NewRandomUser(), datagenerator.NewRandomUser(), datagenerator.NewRandomUser()}
			role1 := datagenerator.NewRandomRole()

			listRequestParams := models.UserListRequestParams{
				Pagination: models.Pagination{
					PageSize: pointers.ToPtr[uint](5),
					Page:     pointers.ToPtr[uint](1),
				},
				Filter: models.ListFilter{
					Search: "email@email.com",
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

			tokenString, err := services.GenerateTokenStringFromUser(users[0], []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, users[0].ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectUsers(mock, users, listRequestParams)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(BeEmpty())
			Expect(testResponse.Data.Users.List.Data).To(HaveLen(len(users)))
			for i := range testResponse.Data.Users.List.Data {
				Expect(testResponse.Data.Users.List.Data[i].ID).To(Equal(users[i].ID))
				Expect(testResponse.Data.Users.List.Data[i].Email).To(Equal(users[i].Email))
				Expect(testResponse.Data.Users.List.Data[i].FirstName).To(Equal(users[i].FirstName))
				Expect(testResponse.Data.Users.List.Data[i].LastName).To(Equal(users[i].LastName))
			}

			Expect(mock.ExpectationsWereMet()).To(BeNil())
		})

		It("Success withouth page and pageSize params", func() {
			users := []models.User{datagenerator.NewRandomUser(), datagenerator.NewRandomUser(), datagenerator.NewRandomUser()}
			role1 := datagenerator.NewRandomRole()

			listRequestParams := models.UserListRequestParams{
				Filter: models.ListFilter{
					Search: "email@email.com",
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

			tokenString, err := services.GenerateTokenStringFromUser(users[0], []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, users[0].ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectUsersWithoutPaging(mock, users, listRequestParams)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(BeEmpty())
			Expect(testResponse.Data.Users.List.Data).To(HaveLen(len(users)))
			for i := range testResponse.Data.Users.List.Data {
				Expect(testResponse.Data.Users.List.Data[i].ID).To(Equal(users[i].ID))
				Expect(testResponse.Data.Users.List.Data[i].Email).To(Equal(users[i].Email))
				Expect(testResponse.Data.Users.List.Data[i].FirstName).To(Equal(users[i].FirstName))
				Expect(testResponse.Data.Users.List.Data[i].LastName).To(Equal(users[i].LastName))
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

			graphRequest := graphqlRequest{
				Query: query,
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(user, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectGetUserByID(mock, user, user.ID)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Profile.ID).To(Equal(user.ID))
		})
	})

	Context("Create", func() {
		It("Succeeds", func() {
			contextUser := datagenerator.NewRandomUser()
			user := datagenerator.NewRandomUser()
			role1 := datagenerator.NewRandomRole()

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

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.CREATE).Return(true, nil)
			sqlexpectations.ExpectCreateUser(mock, contextUser.ID, user)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Create.Code).To(Equal(resolvers.NewUint(200)))
		})
	})

	Context("Update", func() {
		var graphql string = `mutation {
			users {
				update (id: "%v", model: {
					email: ""
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
			role1 := datagenerator.NewRandomRole()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(graphql, user.ID, user.FirstName, user.LastName),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.UPDATE).Return(true, nil)
			sqlexpectations.ExpectUpdateUser(mock, contextUser.ID, user)

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
			role1 := datagenerator.NewRandomRole()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(graphql, user.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlUserResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.DELETE).Return(true, nil)
			sqlexpectations.ExpectDeleteUser(mock, user.ID)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())
			Expect(testResponse.Errors).To(BeEmpty())
			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Users.Delete.Code).To(Equal(resolvers.NewUint(200)))
		})
	})
})
