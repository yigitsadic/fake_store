package main

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func parseUserIdFromJWT(givenToken string) string {
	token, err := jwt.Parse(givenToken, func(tk *jwt.Token) (interface{}, error) {
		return []byte("FAKE_STORE_AUTH"), nil
	})

	if err == nil {
		tokenMap := token.Claims.(jwt.MapClaims)
		if userId, ok := tokenMap["jti"].(string); ok {
			return userId
		}
	}

	return ""
}

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		token := strings.Replace(authHeader, "Bearer ", "", 1)
		userId := parseUserIdFromJWT(token)

		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
