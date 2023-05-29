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
	"github.com/followthepattern/adapticc/pkg/api"
	"github.com/followthepattern/adapticc/pkg/api/graphql/resolvers"
	"github.com/followthepattern/adapticc/pkg/api/middlewares"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/tests/datagenerator"
	"github.com/followthepattern/adapticc/pkg/tests/sqlexpectations"
	"github.com/golang-jwt/jwt/v4"
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
		mdb     *sql.DB
		mock    sqlmock.Sqlmock
		ctx     context.Context
		cfg     config.Config
		cont    *container.Container
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
		if err != nil {
			panic(err)
		}
		cont, err = NewMockedContainer(ctx, mdb, cfg)
		if err != nil {
			panic(err)
		}
		handler, err = api.GetRouter(cont)
		if err != nil {
			panic(err)
		}

	})

	Context("Single", func() {
		var query string = `
		{
			products {
				single (productID: "%s") {
					productID
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
				Query: fmt.Sprintf(query, *product.ProductID),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, tokenString)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(*testResponse.Data.Products.Single.ProductID).To(Equal(*product.ProductID))
			Expect(*testResponse.Data.Products.Single.Title).To(Equal(*product.Title))
			Expect(*testResponse.Data.Products.Single.Description).To(Equal(*product.Description))
		})
	})

	Context("List", func() {
		var query string = `
		query {
			products {
				list (filter:{
					productID: "%v",
					pageSize: %v,
					page: %v,
				}
				) {
					count
					data {
						productID
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

			sqlexpectations.ExpectProducts(mock, *user.ID, page, pageSize, []models.Product{product})

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
				Query: fmt.Sprintf(query, *product.ProductID, pageSize, page),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, tokenString)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(*testResponse.Data.Products.List.Data[0].ProductID).To(Equal(*product.ProductID))
			Expect(*testResponse.Data.Products.List.Data[0].Title).To(Equal(*product.Title))
			Expect(*testResponse.Data.Products.List.Data[0].Description).To(Equal(*product.Description))
		})
	})

	Context("Create", func() {
		var query string = `
		mutation {
			products {
				create (title: "%v", description: "%v") {
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
			httpRequest.Header.Set(middlewares.AuthorizationHeader, tokenString)

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
				update (productID: "%v", title: "%v", description: "%v") {
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
				Query: fmt.Sprintf(query, *product.ProductID, *product.Title, *product.Description),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, tokenString)

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
				delete (productID: "%v") {
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
				Query: fmt.Sprintf(query, *product.ProductID),
			}

			request, _ := json.Marshal(graphRequest)

			testResponse := &graphqlProductResponse{}

			httpRequest := httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request))
			httpRequest.Header.Set(middlewares.AuthorizationHeader, tokenString)

			code, err := runRequest(handler, httpRequest, testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(BeEmpty())

			Expect(code).To(Equal(http.StatusOK))
			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(testResponse.Data.Products.Delete.Code.Value()).To(Equal(uint(http.StatusOK)))
		})
	})
})
