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

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/mocks"
	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/api/middlewares"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/tests/datagenerator"
	"github.com/followthepattern/adapticc/pkg/tests/sqlexpectations"
	"github.com/followthepattern/graphql-go/errors"
	"github.com/golang-jwt/jwt/v4"
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
		mdb      *sql.DB
		mock     sqlmock.Sqlmock
		ctx      context.Context
		cfg      config.Config
		handler  http.Handler
		mockCtrl *gomock.Controller
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
		ac := accesscontrol.Config{
			Kind:   "product",
			Cerbos: mocks.NewMockClient(mockCtrl),
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

		It("Success", func() {
			user := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()

			sqlexpectations.ExpectProduct(mock, *user.ID, product)

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
				Query: fmt.Sprintf(query, *product.ID),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

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
			user := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()

			page := 2
			pageSize := 10

			filter := models.ListFilter{
				Search: product.ID,
			}

			sqlexpectations.ExpectProducts(mock, *user.ID, filter, page, pageSize, []models.Product{product})

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
				Query: fmt.Sprintf(query, pageSize, page, *product.ID),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

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
			user := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()

			expiresAt := time.Now().Add(time.Hour * 24)

			sqlexpectations.CreateProduct(mock, *user.ID, product)

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
				Query: fmt.Sprintf(query, *product.Title, *product.Description),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

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
			user := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()

			sqlexpectations.UpdateProduct(mock, *user.ID, product)

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
				Query: fmt.Sprintf(query, *product.ID, *product.Title, *product.Description),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

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
			user := datagenerator.NewRandomUser()
			product := datagenerator.NewRandomProduct()

			sqlexpectations.DeleteProduct(mock, *user.ID, product)

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
				Query: fmt.Sprintf(query, *product.ID),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, fmt.Sprintf("Bearer %s", tokenString))

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Delete.Code.Value()).To(Equal(uint(http.StatusOK)))
		})
	})
})
