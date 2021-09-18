package auth

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthenticated(t *testing.T) {
	t.Run("it should return error if not found on context", func(t *testing.T) {
		ctx := context.Background()

		got, err := Authenticated(ctx)

		assert.NotNil(t, err)
		assert.Equal(t, "", got)
	})

	t.Run("it should return user id from context", func(t *testing.T) {
		ctx := context.Background()
		expected := "Hello"

		ctx = context.WithValue(ctx, userIDCtxKey, expected)

		got, err := Authenticated(ctx)

		assert.Nil(t, err)
		assert.Equal(t, expected, got)
	})
}
