package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/container"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/golang-jwt/jwt/v4"

	"go.uber.org/zap"
)

type JWT struct {
	logger *zap.Logger
	cfg    config.Config
}

func NewJWT(cont *container.Container) JWT {
	result := JWT{
		logger: cont.GetLogger(),
		cfg:    cont.GetConfig(),
	}
	return result
}

func (a JWT) Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userContext := &models.Guest
		tokenString := r.Header.Get(AuthorizationHeader)

		tokenString = getToken(tokenString)

		if tokenString != "" {
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(a.cfg.Server.HmacSecret), nil
			})

			if err == nil {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					user, err := getUserContextFromClaims(claims)
					if err == nil {
						userContext = user
					} else {
						a.logger.Error(err.Error())
					}
				}
			} else {
				a.logger.Error(err.Error())
			}
		}
		ctx := context.WithValue(r.Context(), utils.CtxUserKey, userContext)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func getToken(tokenString string) string {
	tokens := strings.Split(tokenString, " ")

	if len(tokens) < 2 {
		return ""
	}

	return tokens[1]
}
