package test_api

import (
	"bytes"
	"context"
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
	"github.com/followthepattern/graphql-go/errors"
	"github.com/golang/mock/gomock"
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
		ctx        context.Context
		cfg        config.Config
		handler    http.Handler
		mockCtrl   *gomock.Controller
		mockCerbos *mocks.MockClient
	)

	BeforeEach(func() {
		ctx = context.Background()
		cfg = config.Config{
			Server: config.Server{
				HmacSecret: "test",
			}}
		var err error

		mdb, mock, err = sqlmock.New()
		if err != nil {
			panic(err)
		}

		mockCtrl = gomock.NewController(GinkgoT())
		mockCerbos = mocks.NewMockClient(mockCtrl)
		ac := accesscontrol.Config{
			Kind:   "product",
			Cerbos: mockCerbos,
		}.Build()

		handler = NewMockHandler(ctx, ac, mdb, cfg)

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
				Query: fmt.Sprintf(query, *product.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, *contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectProduct(mock, product)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(*testResponse.Data.Products.Single.ID).To(Equal(*product.ID))
			Expect(*testResponse.Data.Products.Single.Title).To(Equal(*product.Title))
			Expect(*testResponse.Data.Products.Single.Description).To(Equal(*product.Description))
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

			page := 2
			pageSize := 10

			filter := models.ListFilter{
				Search: product.ID,
			}

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, pageSize, page, *product.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, *contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectProducts(mock, filter, page, pageSize, []models.Product{product})

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(*testResponse.Data.Products.List.Data[0].ID).To(Equal(*product.ID))
			Expect(*testResponse.Data.Products.List.Data[0].Title).To(Equal(*product.Title))
			Expect(*testResponse.Data.Products.List.Data[0].Description).To(Equal(*product.Description))
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
				Query: fmt.Sprintf(query, *product.Title, *product.Description),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, *contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.CREATE).Return(true, nil)
			sqlexpectations.CreateProduct(mock, *contextUser.ID, product)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Create.Code.Value()).To(Equal(uint(http.StatusOK)))
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
				Query: fmt.Sprintf(query, *product.ID, *product.Title, *product.Description),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, *contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.UPDATE).Return(true, nil)
			sqlexpectations.UpdateProduct(mock, *contextUser.ID, product)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Update.Code.Value()).To(Equal(uint(http.StatusOK)))
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
				Query: fmt.Sprintf(query, *product.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := services.GenerateTokenStringFromUser(contextUser, []byte(cfg.Server.HmacSecret))
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRolesByUserID(mock, []models.Role{role1}, *contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.DELETE).Return(true, nil)
			sqlexpectations.DeleteProduct(mock, product)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Delete.Code.Value()).To(Equal(uint(http.StatusOK)))
		})
	})
})
