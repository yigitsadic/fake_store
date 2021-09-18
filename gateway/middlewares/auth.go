package middlewares

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

type ctxKey string

const (
	userIDCtxKey = ctxKey("userID")
)

var (
	secret = []byte("FAKE_STORE_AUTH")
)

func parseUserIDFromJWT(givenToken string) string {
	token, err := jwt.Parse(givenToken, func(tk *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err == nil {
		tokenMap := token.Claims.(jwt.MapClaims)
		if userID, ok := tokenMap["jti"].(string); ok {
			return userID
		}
	}

	return ""
}

// Auth chi http middleware for reading token in header['authorization']
func Auth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		token := strings.Replace(authHeader, "Bearer ", "", 1)
		userID := parseUserIDFromJWT(token)

		ctx := r.Context()
		ctx = context.WithValue(ctx, userIDCtxKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
