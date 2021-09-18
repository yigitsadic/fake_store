package utilities

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_GenerateJWTToken(t *testing.T) {
	t.Run("it creates jwt token successfully", func(t *testing.T) {
		s := "ABCDEF"
		c := map[string]string{
			"avatar":   "https://via.placeholder.com/150",
			"fullName": "John Doe",
			"id":       "abcd-efg-12-llu",
		}

		generated := GenerateJWTToken(s, c["id"], c["avatar"], c["fullName"])

		assert.True(t, strings.Contains(generated, "ey"))

		cl := claims{}

		_, err := jwt.ParseWithClaims(generated, &cl, func(token *jwt.Token) (interface{}, error) {
			return []byte(s), nil
		})

		assert.Nil(t, err)
		assert.Equal(t, c["avatar"], cl.Avatar)
		assert.Equal(t, c["fullName"], cl.FullName)
		assert.Equal(t, c["id"], cl.Id)
	})
}
