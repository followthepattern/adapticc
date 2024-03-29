package test_api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"

	"log/slog"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/followthepattern/adapticc/accesscontrol"
	"github.com/followthepattern/adapticc/api"
	"github.com/followthepattern/adapticc/api/graphql"
	"github.com/followthepattern/adapticc/api/rest"
	"github.com/followthepattern/adapticc/config"
	"github.com/followthepattern/adapticc/container"
	"github.com/followthepattern/adapticc/controllers"
	"github.com/followthepattern/adapticc/features/mail"

	"github.com/golang-jwt/jwt/v4"
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

func NewMockHandler(ac accesscontrol.AccessControl, emailClient mail.Email, db *sql.DB, cfg config.Config, jwtKeys config.JwtKeyPair) http.Handler {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cont := container.New(ac, emailClient, db, cfg, logger, jwtKeys)

	ctrls := controllers.New(cont)

	graphqlHandler := graphql.NewHandler(ctrls)

	restHandler := rest.New(ctrls)

	return api.NewHttpApi(cfg, jwtKeys, graphqlHandler, restHandler, logger)
}

var replaceEmptySpacesToSpace = regexp.MustCompile(`\s+`)
var removeEmptySpaceBeforeSelect = regexp.MustCompile(`\(\s+SELECT`)
var replaceOpeningBracket = regexp.MustCompile(`\(`)
var replaceClosingBracket = regexp.MustCompile(`\)`)
var replacePipe = regexp.MustCompile(`\|`)

func stripQuery(q string) (s string) {
	return strings.TrimSpace(replaceEmptySpacesToSpace.ReplaceAllString(q, " "))
}

func getMockJWTKeys() (config.JwtKeyPair, error) {
	publicKeystring := []byte(`-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEA5xdfauTdmBzpx9McaSvMaZRhUlGlcrTAe0IHQqnwMjs=
-----END PUBLIC KEY-----
`)

	ed25519PublicKey, err := jwt.ParseEdPublicKeyFromPEM(publicKeystring)
	if err != nil {
		return config.JwtKeyPair{}, err
	}

	privateKeystring := []byte(`-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEIB8klgJvgq1xk5AMw4lZFfQvjCwaRxH8ghJHGn8KQtUM
-----END PRIVATE KEY-----`)

	ed25519PrivateKey, err := jwt.ParseEdPrivateKeyFromPEM(privateKeystring)
	if err != nil {
		return config.JwtKeyPair{}, err
	}

	return config.JwtKeyPair{
		Public:  ed25519PublicKey,
		Private: ed25519PrivateKey,
	}, nil
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
