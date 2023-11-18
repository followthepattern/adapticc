package test_api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/mocks"
	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/api/middlewares"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/services"
	"github.com/followthepattern/adapticc/pkg/tests/datagenerator"
	"github.com/followthepattern/adapticc/pkg/tests/sqlexpectations"
	"github.com/golang/mock/gomock"
	"github.com/graph-gophers/graphql-go/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type graphqlProductResponse struct {
	Data   productData          `json:"data"`
	Errors []*errors.QueryError `json:"errors,omitempty"`
}

type productData struct {
	Products products `json:"products"`
}

type products struct {
	Single models.Product                      `json:"single,omitempty"`
	List   models.ListResponse[models.Product] `json:"list,omitempty"`
	Create resolvers.ResponseStatus            `json:"create,omitempty"`
	Update resolvers.ResponseStatus            `json:"update,omitempty"`
	Delete resolvers.ResponseStatus            `json:"delete,omitempty"`
}

var _ = Describe("Product Test", func() {
	var (
		mdb        *sql.DB
		mock       sqlmock.Sqlmock
		cfg        config.Config
		jwtKeys    config.JwtKeyPair
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
		Expect(err).ShouldNot(HaveOccurred())

		jwtKeys, err = getMockJWTKeys()
		Expect(err).ShouldNot(HaveOccurred())

		mockCtrl = gomock.NewController(GinkgoT())
		mockCerbos = mocks.NewMockClient(mockCtrl)
		ac := accesscontrol.Config{
			Kind:   "product",
			Cerbos: mockCerbos,
		}.Build()

		handler = NewMockHandler(ac, nil, mdb, cfg, jwtKeys)

	})

	Context("Single", func() {
		var query string = `
		{
			products {
				single (id: "%s") {
					id
					title
					description
				}
			}
		}`

		It("Succeeds", func() {
			contextUser := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()
			role1 := datagenerator.NewRandomRole()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, product.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectProduct(mock, product)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Single.ID.Data).To(Equal(product.ID.Data))
			Expect(testResponse.Data.Products.Single.Title.Data).To(Equal(product.Title.Data))
			Expect(testResponse.Data.Products.Single.Description.Data).To(Equal(product.Description.Data))
		})
	})

	Context("List", func() {
		var query string = `query {
			products {
				list (
					pagination: { pageSize: %v, page: %v }
					filter: { search: "%s" }
				) {
					count
					data {
						id
						title
						description
					}
					page
					pageSize
				}
			}
		}`

		It("Success", func() {
			product := datagenerator.NewRandomProduct()
			contextUser := datagenerator.NewRandomUser()
			role1 := datagenerator.NewRandomRole()

			page := 4
			pageSize := 10

			filter := models.ListFilter{
				Search: product.ID.Data,
			}

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, pageSize, page, product.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectProducts(mock, filter, page, pageSize, []models.Product{product})

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.List.Data[0].ID.Data).To(Equal(product.ID.Data))
			Expect(testResponse.Data.Products.List.Data[0].Title.Data).To(Equal(product.Title.Data))
			Expect(testResponse.Data.Products.List.Data[0].Description.Data).To(Equal(product.Description.Data))
		})
	})

	Context("Create", func() {
		var query string = `
		mutation {
			products {
				create(model: {
					title: "%s",
					description: "%s"
				}) {
					code
				}
			}
		}`

		It("Success", func() {
			contextUser := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()
			role1 := datagenerator.NewRandomRole()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, product.Title, product.Description),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.CREATE).Return(true, nil)
			sqlexpectations.CreateProduct(mock, contextUser.ID, product)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Create.Code).To(Equal(int32(http.StatusCreated)))
		})
	})

	Context("Update", func() {
		var query string = `
		mutation {
			products {
				update(id: "%s",
					model: {
					title: "%s",
					description: "%s"
				}) {
					code
				}
			}
		}`

		It("Success", func() {
			contextUser := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()
			role1 := datagenerator.NewRandomRole()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, product.ID, product.Title, product.Description),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.UPDATE).Return(true, nil)
			sqlexpectations.UpdateProduct(mock, contextUser.ID, product)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Update.Code).To(Equal(int32(http.StatusOK)))
		})
	})

	Context("Delete", func() {
		var query string = `
		mutation {
			products {
				delete (id: "%v") {
					code
				}
			}
		}`

		It("Success", func() {
			contextUser := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()
			role1 := datagenerator.NewRandomRole()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, product.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.DELETE).Return(true, nil)
			sqlexpectations.DeleteProduct(mock, product)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Delete.Code).To(Equal(int32(http.StatusOK)))
		})
	})
})
