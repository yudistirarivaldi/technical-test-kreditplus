package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/yudistirarivaldi/technical-test-kreditplus/internal/utils"
)

type contextKey string

const ConsumerIDKey contextKey = "consumer_id"

func JWTMiddleware(secret string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"responseCode": "01",
				"message":      "Missing or invalid Authorization header",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		consumerID, err := utils.ParseJWT(tokenStr, secret)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, map[string]string{
				"responseCode": "01",
				"message":      "Invalid or expired token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), ConsumerIDKey, consumerID)

		next(w, r.WithContext(ctx))
	}
}

func GetConsumerIDFromContext(ctx context.Context) (int64, bool) {
	consumerID, ok := ctx.Value(ConsumerIDKey).(int64)
	return consumerID, ok
}
