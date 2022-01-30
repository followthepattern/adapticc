package api

import (
	"backend/internal"
	"backend/internal/container"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const url = "/graphql"

type graphqlRequest struct {
	Query string `json:"query"`
}

func runRequest(t *testing.T, srv http.Handler, r *http.Request, data interface{}) (int, error) {
	response := httptest.NewRecorder()
	srv.ServeHTTP(response, r)

	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(data)
	if err != nil {
		return 0, fmt.Errorf("couldn't decode Response json: %v", err)
	}

	return response.Code, nil
}

func NewMockedContainer(db *sql.DB) (container.IContainer, error) {
	ctx := context.Background()

	cont := container.New(
		&ctx,
		nil,
		db,
		nil)

	err := internal.RegisterDependencies(cont)
	if err != nil {
		return nil, err
	}

	return cont, nil
}
