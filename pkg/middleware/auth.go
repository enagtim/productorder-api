package middleware

import (
	"context"
	"net/http"
	"order-api/configs"
	"order-api/pkg/jwt"
	"order-api/pkg/messages"
	"strings"
)

type key string

const (
	ContextUserID key = "ContextUserID"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			messages.SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(header, "Bearer ")
		isValid, data := jwt.NewJWT(config.Auth.Secret).ParseToken(token)
		if !isValid {
			messages.SendJSONError(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserID, data.UserID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
