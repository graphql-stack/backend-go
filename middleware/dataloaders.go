package middleware

import (
	"context"
	"github.com/graphql-stack/backend-go/dataloader"
	"net/http"
	"time"
)

const userloaderContextKey = "context_user_loader"

func DataloadersMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userLoader := dataloader.NewUserLoader(100, 2*time.Millisecond)

			// put it in context
			ctx := context.WithValue(r.Context(), userloaderContextKey, userLoader)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetUserLoader(ctx context.Context) *dataloader.UserLoader {
	return ctx.Value(userloaderContextKey).(*dataloader.UserLoader)
}
