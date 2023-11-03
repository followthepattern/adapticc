package middlewares

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/followthepattern/adapticc/pkg/config"
	"github.com/followthepattern/adapticc/pkg/models"
	"github.com/followthepattern/adapticc/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	logger *slog.Logger
	cfg    config.Config
}

func NewJWT(logger *slog.Logger, cfg config.Config) JWT {
	result := JWT{
		logger: logger,
		cfg:    cfg,
	}
	return result
}

func (a JWT) Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var userContext models.User

		defer func() {
			ctx := context.WithValue(r.Context(), utils.CtxUserKey, userContext)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}()

		headerValue := r.Header.Get(AuthorizationHeader)
		tokenString := getToken(headerValue)

		if tokenString == "" {
			return
		}

		token, err := jwt.Parse(tokenString, a.keyFunc)
		if err != nil {
			a.logger.Error(err.Error())
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			a.logger.Error(InvalidToken)
			return
		}

		user, err := getAuthorizedUserFromClaims(claims)
		if err != nil {
			a.logger.Error(err.Error())
			return
		}

		userContext = *user
	}
	return http.HandlerFunc(fn)
}

func (a JWT) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(a.cfg.Server.HmacSecret), nil
}
