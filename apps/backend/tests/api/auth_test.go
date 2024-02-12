package test_api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/smtp"

	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/features/auth"
	"github.com/followthepattern/adapticc/mocks"
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

type graphqlAuthResponse struct {
	Data   authenticationData `json:"data"`
	Errors []*errors.QueryError
}

type authenticationData struct {
	Authentication authentication `json:"authentication"`
}

type authentication struct {
	Register auth.RegisterResponse `json:"register"`
	Login    auth.LoginResponse    `json:"login"`
}

var _ = Describe("Authentication", func() {
	var (
		mdb     *sql.DB
		mock    sqlmock.Sqlmock
		cfg     config.Config
		jwtKeys config.JwtKeyPair
		handler http.Handler

		mockCtrl  *gomock.Controller
		mockEmail *mocks.MockEmail

		testResponse  *graphqlAuthResponse
		generatedUser auth.AuthUser
		password      string
	)

	BeforeEach(func() {
		var err error
		mdb, mock, err = sqlmock.New()
		Expect(err).To(BeNil())
		cfg = config.Config{
			Server: config.Server{
				HmacSecret:            "test",
				GraphqlSchemaFilepath: "./../../api/graphql/schema/schema.graph",
			},
			Mail: config.Mail{
				Addr:     "addr",
				Host:     "host",
				Username: "test-username",
				Password: "test-password",
			},
		}
		jwtKeys, err = getMockJWTKeys()
		Expect(err).ShouldNot(HaveOccurred())

		ac := accesscontrol.Config{}.Build()
		mockCtrl = gomock.NewController(GinkgoT())
		mockEmail = mocks.NewMockEmail(mockCtrl)

		handler = NewMockHandler(ac, mockEmail, mdb, cfg, jwtKeys)

		testResponse = &graphqlAuthResponse{}
		password = datagenerator.String(13)
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		Expect(err).Should(BeNil())

		generatedUser = datagenerator.NewRandomAuthUser(passwordHash)
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
			sqlexpectations.ExpectGetAuthUserByEmail(mock, generatedUser, generatedUser.Email)

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, generatedUser.Email, password),
			}

			request, _ := json.Marshal(graphRequest)

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(HaveLen(0))

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))
		})

		It("Wrong password", func() {
			sqlexpectations.ExpectGetAuthUserByEmail(mock, generatedUser, generatedUser.Email)

			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, generatedUser.Email, "wrong-password"),
			}

			request, _ := json.Marshal(graphRequest)

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(testResponse.Errors).To(HaveLen(1))
			Expect(testResponse.Errors[0].Message).To(Equal(auth.WRONG_EMAIL_OR_PASSWORD))

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
				Query: fmt.Sprintf(queryTemplate, generatedUser.Email, generatedUser.FirstName, generatedUser.LastName, password),
			}
			request, _ := json.Marshal(graphRequest)

			sqlexpectations.ExpectVerifyEmail(mock, 0, generatedUser.Email)

			sqlexpectations.ExpectCreateAuthUser(mock, generatedUser)

			mailTemplate := auth.GetActivationMailTemplate(cfg, types.StringFrom(""), generatedUser.Email)

			mockEmail.EXPECT().SetFrom(gomock.Any())
			mockEmail.EXPECT().SetTo(mailTemplate.To)
			mockEmail.EXPECT().SetSubject(mailTemplate.Subject)
			mockEmail.EXPECT().SetText(gomock.Any())
			mockEmail.EXPECT().SetHTML(gomock.Any())
			mockEmail.EXPECT().Send(cfg.Mail.Addr,
				smtp.PlainAuth(
					"",
					cfg.Mail.Username,
					cfg.Mail.Password,
					cfg.Mail.Host))

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(HaveLen(0))

			Expect(generatedUser.FirstName.Data).To(Equal(testResponse.Data.Authentication.Register.FirstName.Data))
			Expect(generatedUser.LastName.Data).To(Equal(testResponse.Data.Authentication.Register.LastName.Data))
			Expect(generatedUser.Email.Data).To(Equal(testResponse.Data.Authentication.Register.Email.Data))

			Expect(mock.ExpectationsWereMet()).To(BeNil())
		})

		It("Duplicated email", func() {
			graphRequest := graphqlRequest{
				Query: fmt.Sprintf(queryTemplate, generatedUser.Email, generatedUser.FirstName, generatedUser.LastName, password),
			}
			request, _ := json.Marshal(graphRequest)

			sqlexpectations.ExpectVerifyEmail(mock, 1, generatedUser.Email)

			code, err := runRequest(handler, httptest.NewRequest("POST", graphqlURL, bytes.NewReader(request)), testResponse)
			Expect(err).To(BeNil())

			Expect(mock.ExpectationsWereMet()).To(BeNil())
			Expect(code).To(Equal(http.StatusOK))

			Expect(testResponse.Errors).To(HaveLen(1))
			Expect(testResponse.Errors[0].Message).To(Equal(fmt.Sprintf(auth.EMAIL_IS_ALREADY_IN_USE_PATTERN, generatedUser.Email)))
		})
	})
})
