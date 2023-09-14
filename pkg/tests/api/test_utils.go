package test_api

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/pkg/api"
	"github.com/followthepattern/adapticc/pkg/api/graphql"
	"github.com/followthepattern/adapticc/pkg/api/rest"
	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/controllers"
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

func NewMockHandler(ctx context.Context, db *sql.DB, cfg config.Config) http.Handler {
	logger := zap.NewExample()

	ctrls := controllers.New(ctx, db, cfg, logger)

	graphqlHandler := graphql.New(ctrls)

	restHandler := rest.New(ctrls)

	return api.NewHttpApi(cfg, graphqlHandler, restHandler, logger)
}

var replaceEmptySpacesToSpace = regexp.MustCompile(`\s+`)
var removeEmptySpaceBeforeSelect = regexp.MustCompile(`\(\s+SELECT`)
var replaceOpeningBracket = regexp.MustCompile(`\(`)
var replaceClosingBracket = regexp.MustCompile(`\)`)
var replacePipe = regexp.MustCompile(`\|`)

func stripQuery(q string) (s string) {
	return strings.TrimSpace(replaceEmptySpacesToSpace.ReplaceAllString(q, " "))
}

// strip out new lines and trim spaces
func escapeRegexpCharacters(q string) string {
	b := replaceEmptySpacesToSpace.ReplaceAllString(q, " ")
	b = replaceOpeningBracket.ReplaceAllString(b, "\\(")
	b = replaceClosingBracket.ReplaceAllString(b, "\\)")
	b = replacePipe.ReplaceAllString(b, "\\|")
	b = removeEmptySpaceBeforeSelect.ReplaceAllString(b, "(SELECT")

	return strings.TrimSpace(b)
}

var sqlCompareFunc = sqlmock.QueryMatcherFunc(func(expectedRegex, actualSQL string) error {
	exprx := escapeRegexpCharacters(expectedRegex)
	actual := stripQuery(actualSQL)
	re, err := regexp.Compile(exprx)
	if err != nil {
		return err
	}

	locations := re.FindStringIndex(actual)

	if locations == nil {
		return fmt.Errorf(`actual sql doesn't match:
		%s
		with regexp:
		%s`, actual, exprx)
	}

	return nil
})
