package api

import (
	"backend/internal/api/graphql_api"
	"backend/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/stretchr/testify/assert"
)

type graphqlUserResponse struct {
	Data   userData             `json:"data"`
	Errors []*errors.QueryError `json:"errors,omitempty"`
}

type userData struct {
	Users users `json:"users"`
}

type users struct {
	Single models.User `json:"single,omitempty"`
}

func TestGetUser(t *testing.T) {
	mdb, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	url := "/graphql"
	cont, err := NewMockedContainer(mdb)

	if err != nil {
		t.Fatal(err)
	}

	handler, err := graphql_api.NewHandler(cont)

	if err != nil {
		t.Fatal(err)
	}

	queryTemplate := `query {
		users {
			single(id: "%v") {
				id
				email
			}
		}
	}
	`
	id := uuid.New()
	query := fmt.Sprintf(queryTemplate, id.String())

	graphRequest := graphqlRequest{
		Query: query,
	}

	request, _ := json.Marshal(graphRequest)

	sqlQuery := fmt.Sprintf(`SELECT "active", "created_at", "creation_user_id", "email", "first_name", "id", "last_login_at", "last_name", "password_hash", "salt", "update_user_id", "updated_at" FROM "users" WHERE \("id" = '%v'\) LIMIT 1`, id.String())

	mock.ExpectQuery(sqlQuery).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(id.String()))

	testResponse := &graphqlUserResponse{}

	t.Run("Users single item", func(t *testing.T) {
		code, err := runRequest(t, handler, httptest.NewRequest("POST", url, bytes.NewReader(request)), testResponse)
		assert.Nil(t, err)

		if len(testResponse.Errors) > 0 {
			assert.Nil(t, testResponse.Errors[0].Message)
		}

		assert.Nil(t, mock.ExpectationsWereMet())
		assert.Equal(t, code, http.StatusOK)
		assert.Equal(t, id.String(), *testResponse.Data.Users.Single.ID)
	})
}
