package test_api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/api/middlewares"
	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/features/product"
	"github.com/followthepattern/adapticc/mocks"
	"github.com/followthepattern/adapticc/models"
	"github.com/followthepattern/adapticc/tests/datagenerator"
	"github.com/followthepattern/adapticc/tests/sqlexpectations"
	"github.com/golang/mock/gomock"
	"github.com/graph-gophers/graphql-go/errors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/bcrypt"
)

type graphqlProductResponse struct {
	Data   productData          `json:"data"`
	Errors []*errors.QueryError `json:"errors,omitempty"`
}

type productData struct {
	Products products `json:"products"`
}

type products struct {
	Single product.ProductModel                      `json:"single,omitempty"`
	List   models.ListResponse[product.ProductModel] `json:"list,omitempty"`
	Create models.ResponseStatus                     `json:"create,omitempty"`
	Update models.ResponseStatus                     `json:"update,omitempty"`
	Delete models.ResponseStatus                     `json:"delete,omitempty"`
}

var _ = Describe("Product Test", func() {
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

		password := datagenerator.String(18)
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		Expect(err).Should(BeNil())
		contextUser = datagenerator.NewRandomAuthUser(passwordHash)

		roleIDs = []string{datagenerator.String(18), datagenerator.String(18)}
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
			product1 := datagenerator.NewRandomProduct()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, product1.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectProduct(mock, product1)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Single.ID.Data).To(Equal(product1.ID.Data))
			Expect(testResponse.Data.Products.Single.Title.Data).To(Equal(product1.Title.Data))
			Expect(testResponse.Data.Products.Single.Description.Data).To(Equal(product1.Description.Data))
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
			product1 := datagenerator.NewRandomProduct()

			page := 4
			pageSize := 10

			filter := models.ListFilter{
				Search: product1.ID,
			}

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, pageSize, page, product1.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.READ).Return(true, nil)
			sqlexpectations.ExpectProducts(mock, filter, page, pageSize, []product.ProductModel{product1})

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.List.Data[0].ID.Data).To(Equal(product1.ID.Data))
			Expect(testResponse.Data.Products.List.Data[0].Title.Data).To(Equal(product1.Title.Data))
			Expect(testResponse.Data.Products.List.Data[0].Description.Data).To(Equal(product1.Description.Data))
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
			product1 := datagenerator.NewRandomProduct()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, product1.Title, product1.Description),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.CREATE).Return(true, nil)
			sqlexpectations.CreateProduct(mock, contextUser.ID, product1)

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
			product1 := datagenerator.NewRandomProduct()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, product1.ID, product1.Title, product1.Description),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.UPDATE).Return(true, nil)
			sqlexpectations.UpdateProduct(mock, contextUser.ID, product1)

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
			product1 := datagenerator.NewRandomProduct()

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(query, product1.ID),
			}
			request, _ := json.Marshal(graphRequest)

			tokenString, err := auth.GenerateTokenStringFromUser(contextUser, jwtKeys.Private)
			Expect(err).To(BeNil())

			testResponse := &graphqlProductResponse{}
			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("%s %s", middlewares.BearerPrefix, tokenString))

			sqlexpectations.ExpectRoleIDsByUserID(mock, roleIDs, contextUser.ID)
			mockCerbos.EXPECT().IsAllowed(gomock.Any(), gomock.Any(), gomock.Any(), accesscontrol.DELETE).Return(true, nil)
			sqlexpectations.DeleteProduct(mock, product1)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Delete.Code).To(Equal(int32(http.StatusOK)))
		})
	})
})
