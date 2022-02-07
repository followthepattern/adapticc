package middlewares

import (
	"backend/internal/container"
	"backend/internal/services"
	"backend/internal/utils"
	"fmt"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer"
)

var CtxUserKey = &utils.ContextKey{Name: "ctx-user"}

type Auth struct {
	logger *zap.Logger
	us     services.User
}

func NewAuth(cont container.IContainer) Auth {
	result := Auth{
		logger: cont.GetLogger(),
	}
	cont.Resolve2(&result.us)
	return result
}

func (a Auth) Authenticate(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(AuthorizationHeader)
		if authHeader != "" {
			fmt.Println(authHeader)
			tokenHeader := strings.Split(authHeader, BearerPrefix)
			if len(tokenHeader) == 2 {
				// token := strings.Trim(tokenHeader[1], " ")
				// user, err := a.us.GetByToken(token)
				// if err != nil {
				// 	a.logger.Error(err.Error())
				// }
				// if user != nil && user.ID != nil {
				// 	ctx := context.WithValue(r.Context(), CtxUserKey, models.User{})
				// 	r = r.WithContext(ctx)
				// }
			}
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
