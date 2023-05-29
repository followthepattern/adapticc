package test_api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	internal "github.com/followthepattern/adapticc/pkg"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"go.uber.org/zap"
)

const graphqlURL = "/graphql"

type graphqlRequest struct {
	Query string `json:"query"`
}

func runRequest(srv http.Handler, r *http.Request, data interface{}) (int, error) {
	response := httptest.NewRecorder()
	srv.ServeHTTP(response, r)

	decoder := json.NewDecoder(response.Body)
	err := decoder.Decode(data)
	if err != nil {
		return 0, fmt.Errorf("couldn't decode Response json: %v", err)
	}

	return response.Code, nil
}

func NewMockedContainer(ctx context.Context, db *sql.DB, cfg config.Config) (*container.Container, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	cont := container.New(
		ctx,
		cfg,
		db,
		logger)

	err = internal.RegisterDependencies(cont)
	if err != nil {
		return nil, err
	}

	return cont, nil
}
