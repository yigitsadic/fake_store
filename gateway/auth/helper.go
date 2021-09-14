package auth

import (
	"context"
	"errors"
)

var (
	errUnauthorized = errors.New("unauthorized to access this resource")
)

// Authenticated Tries fetch userID from context parameter and convert it to string.
func Authenticated(ctx context.Context) (string, error) {
	userID := ctx.Value(userIDCtxKey)

	v, ok := userID.(string)
	if !ok {
		return "", errUnauthorized
	}

	if v == "" {
		return "", errUnauthorized
	}

	return v, nil
}
