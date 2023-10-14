package test_api

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/followthepattern/adapticc/pkg/accesscontrol"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/services"
	"github.com/followthepattern/adapticc/pkg/tests/datagenerator"
	"github.com/followthepattern/adapticc/pkg/tests/sqlexpectations"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/graphql-go/errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type graphqlAuthResponse struct {
	Data   authenticationData `json:"data"`
	Errors []*errors.QueryError
}

type authenticationData struct {
	Authentication authentication `json:"authentication"`
}

type authentication struct {
	Register models.RegisterResponse `json:"register"`
	Login    models.LoginResponse    `json:"login"`
}

var _ = Describe("Authentication", func() {
	var (
		mdb     *sql.DB
		mock    sqlmock.Sqlmock
		ctx     context.Context
		cfg     config.Config
		handler http.Handler

		testResponse  *graphqlAuthResponse
		generatedUser models.AuthUser
		password      string
	)

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		mdb, mock, err = sqlmock.New()
		Expect(err).To(BeNil())
		cfg = config.Config{
			Server: config.Server{
				HmacSecret: "test",
			},
		}

		ac := accesscontrol.Config{}.Build()

		handler = NewMockHandler(ctx, ac, mdb, cfg)

		testResponse = &graphqlAuthResponse{}
		password = datagenerator.String(13)
		generatedUser = datagenerator.NewRandomAuthUser(password)
	})

	Context("Login", func() {
		var (
			queryTemplate string = `
			mutation {
				authentication {
					login(email: "%v", password: "%v") {
						jwt
						expires_at
					}
				}
			}`
		)

		It("Success", func() {
			sqlexpectations.ExpectGetAuthUserByEmail(mock, generatedUser, *generatedUser.Email)

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, *generatedUser.Email, password),
			}

			request, _ := json.Marshal(graphRequest)

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(HaveLen(0))

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))
		})

		It("Wrong password", func() {
			sqlexpectations.ExpectGetAuthUserByEmail(mock, generatedUser, *generatedUser.Email)

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, *generatedUser.Email, "wrong-password"),
			}

			request, _ := json.Marshal(graphRequest)

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(HaveLen(1))
			Expect(testResponse.Errors[0].Message).To(Equal(services.WRONG_EMAIL_OR_PASSWORD))

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))
		})
	})

	Context("Register", func() {
		var queryTemplate = `
		mutation {
			authentication {
				register(email: "%v", firstName: "%v", lastName: "%v", password: "%v") {
					email
					first_name
					last_name
				}
			}
		}`

		It("Success", func() {
			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, *generatedUser.Email, *generatedUser.FirstName, *generatedUser.LastName, password),
			}
			request, _ := json.Marshal(graphRequest)

			sqlexpectations.ExpectVerifyEmail(mock, 0, *generatedUser.Email)

			sqlexpectations.ExpectCreateAuthUser(mock, "", generatedUser)

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(HaveLen(0))

			Expect(*generatedUser.FirstName).To(Equal(*testResponse.Data.Authentication.Register.FirstName))
			Expect(*generatedUser.LastName).To(Equal(*testResponse.Data.Authentication.Register.LastName))
			Expect(*generatedUser.Email).To(Equal(*testResponse.Data.Authentication.Register.Email))

			Expect(mock.ExpectationsWereMet()).To(BeNil())
		})

		It("Duplicated email", func() {
			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, *generatedUser.Email, *generatedUser.FirstName, *generatedUser.LastName, password),
			}
			request, _ := json.Marshal(graphRequest)

			sqlexpectations.ExpectVerifyEmail(mock, 1, *generatedUser.Email)

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(HaveLen(1))
			Expect(testResponse.Errors[0].Message).To(Equal(fmt.Sprintf(services.EMAIL_IS_ALREADY_IN_USE_PATTERN, *generatedUser.Email)))
		})
	})
})