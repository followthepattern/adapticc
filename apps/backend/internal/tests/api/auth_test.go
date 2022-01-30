package api

import (
	"backend/internal/api/graphql_api"
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/tests/data_generator"
	"backend/internal/tests/sqlexpectations"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/stretchr/testify/assert"
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

func TestLogin(t *testing.T) {
	mdb, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	cont, err := NewMockedContainer(mdb)

	if err != nil {
		t.Fatal(err)
	}

	handler, err := graphql_api.NewHandler(cont)

	if err != nil {
		t.Fatal(err)
	}

	queryTemplate := `mutation {
		authentication {
			login(email: "%v", password: "%v") {
				jwt
				expires_at
			}
		}
	}
	`

	password := data_generator.String(13)
	resultUser := data_generator.NewRandomUser(password)

	testResponse := &graphqlAuthResponse{}

	t.Run("Success", func(t *testing.T) {
		sqlexpectations.ExpectGetUserByEmail(mock, *resultUser.Email, resultUser)
		sqlexpectations.ExpectInsertToken(mock, resultUser)

		graphRequest := graphqlRequest{
			Query: fmt.Sprintf(queryTemplate, *resultUser.Email, password),
		}

		request, _ := json.Marshal(graphRequest)

		code, err := runRequest(t, handler, httptest.NewRequest("POST", url, bytes.NewReader(request)), testResponse)
		assert.Nil(t, err)

		if len(testResponse.Errors) > 0 {
			assert.Nil(t, testResponse.Errors[0].Message)
		}

		assert.Nil(t, mock.ExpectationsWereMet())
		assert.Equal(t, code, http.StatusOK)
		assert.NotNil(t, testResponse.Data.Authentication.Login.JWT)
	})

	t.Run("Wrong password", func(t *testing.T) {
		sqlexpectations.ExpectGetUserByEmail(mock, *resultUser.Email, resultUser)

		graphRequest := graphqlRequest{
			Query: fmt.Sprintf(queryTemplate, *resultUser.Email, "wrong-password"),
		}

		request, _ := json.Marshal(graphRequest)

		code, err := runRequest(t, handler, httptest.NewRequest("POST", url, bytes.NewReader(request)), testResponse)
		assert.Nil(t, err)

		if len(testResponse.Errors) > 0 {
			assert.Equal(t, services.INCORRECT_EMAIL_OR_PASSWORD, testResponse.Errors[0].Message)
		}

		assert.Nil(t, mock.ExpectationsWereMet())
		assert.Equal(t, code, http.StatusOK)
	})
}

func TestRegister(t *testing.T) {
	mdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	cont, err := NewMockedContainer(mdb)
	if err != nil {
		t.Fatal(err)
	}

	handler, err := graphql_api.NewHandler(cont)
	if err != nil {
		t.Fatal(err)
	}

	password := data_generator.String(10)
	resultUser := data_generator.NewRandomUser(password)

	testResponse := &graphqlAuthResponse{}

	t.Run("Success", func(t *testing.T) {
		queryTemplate := `mutation {
			authentication {
				register(email: "%v", firstName: "%v", lastName: "%v", password: "%v") {
					email
					first_name
					last_name
				}
			}
		}
		`

		graphRequest := graphqlRequest{
			Query: fmt.Sprintf(queryTemplate, *resultUser.Email, *resultUser.FirstName, *resultUser.LastName, password),
		}
		request, _ := json.Marshal(graphRequest)

		sqlQuery := fmt.Sprintf(`INSERT INTO "users" \("active", "created_at", "creation_user_id", "email", "first_name", "id", "last_login_at", "last_name", "password_hash", "salt", "update_user_id", "updated_at"\) VALUES \(TRUE, '.*', NULL, '%v', '%v', '.*', NULL, '%v', '.*', '.*', NULL, NULL\)`,
			*resultUser.Email,
			*resultUser.FirstName,
			*resultUser.LastName)

		sqlexpectations.ExpectGetUserByEmail(mock, *resultUser.Email, models.User{})

		mock.ExpectExec(sqlQuery).
			WillReturnResult(sqlmock.NewResult(1, 1))

		code, err := runRequest(t, handler, httptest.NewRequest("POST", url, bytes.NewReader(request)), testResponse)
		assert.Nil(t, err)

		assert.Nil(t, mock.ExpectationsWereMet())
		assert.Equal(t, code, http.StatusOK)

		if len(testResponse.Errors) == 0 {
			assert.Equal(t, *resultUser.FirstName, *testResponse.Data.Authentication.Register.FirstName)
			assert.Equal(t, *resultUser.LastName, *testResponse.Data.Authentication.Register.LastName)
			assert.Equal(t, *resultUser.Email, *testResponse.Data.Authentication.Register.Email)
		} else {
			assert.Nil(t, testResponse.Errors[0].Message)
		}
	})
	t.Run("Duplicated email", func(t *testing.T) {
		queryTemplate := `mutation {
			authentication {
				register(email: "%v", firstName: "%v", lastName: "%v", password: "%v") {
					email
					first_name
					last_name
				}
			}
		}
		`

		graphRequest := graphqlRequest{
			Query: fmt.Sprintf(queryTemplate, *resultUser.Email, *resultUser.FirstName, *resultUser.LastName, password),
		}
		request, _ := json.Marshal(graphRequest)

		sqlexpectations.ExpectGetUserByEmail(mock, *resultUser.Email, resultUser)

		code, err := runRequest(t, handler, httptest.NewRequest("POST", url, bytes.NewReader(request)), testResponse)
		assert.Nil(t, err)

		assert.Nil(t, mock.ExpectationsWereMet())
		assert.Equal(t, code, http.StatusOK)

		if len(testResponse.Errors) > 0 {
			assert.Equal(t, fmt.Sprintf(services.EMAIL_IS_ALREADY_IN_USE_PATTERN, *resultUser.Email), testResponse.Errors[0].Message)
		}
	})
}
