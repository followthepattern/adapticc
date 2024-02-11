package middlewares

import (
	"context"
	"net/http"

	"github.com/followthepattern/adapticc/features/auth"
	"github.com/google/uuid"
)

var SessionContextHeader = &auth.ContextKey{Name: "session-context-id"}

func SessionContextID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sessionContextID := r.Header.Get(SessionContextHeader.Name)
		if sessionContextID == "" {
			sessionContextID = uuid.New().String()
		}
		ctx = context.WithValue(ctx, SessionContextHeader, sessionContextID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
