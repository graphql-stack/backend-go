package middleware

import (
	"context"
	"github.com/graphql-stack/backend-go/model"
	"github.com/graphql-stack/backend-go/service"
	"net/http"
	"strings"
)

const (
	userContextKey = "context_user"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			a := r.Header.Get("Authorization")

			token := strings.Replace(a, "Bearer ", "", 1)

			if token == "" {
				next.ServeHTTP(w, r)
				return
			}

			u, err := service.GetUserByToken(token)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userContextKey, u)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetCurrentUser(ctx context.Context) *model.User {
	return ctx.Value(userContextKey).(*model.User)
}
