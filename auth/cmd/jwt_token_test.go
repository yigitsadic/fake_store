package main

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
	"testing"
)

func Test_generateJWTToken(t *testing.T) {
	t.Run("it creates jwt token successfully", func(t *testing.T) {
		s := "ABCDEF"
		c := map[string]string{
			"avatar":   "https://via.placeholder.com/150",
			"fullName": "John Doe",
			"id":       "abcd-efg-12-llu",
		}

		generated := generateJWTToken(s, c["id"], c["avatar"], c["fullName"])

		if !strings.Contains(generated, "ey") {
			t.Errorf("expected to produce token with starting ey")
		}

		cl := claims{}

		_, err := jwt.ParseWithClaims(generated, &cl, func(token *jwt.Token) (interface{}, error) {
			return []byte(s), nil
		})
		if err != nil {
			t.Fatalf("unexpected to get an error while parsing token. Err %s", err)
		}

		if cl.Avatar != c["avatar"] {
			t.Errorf("expected avatar was %s but got %s", c["avatar"], cl.Avatar)
		}

		if cl.FullName != c["fullName"] {
			t.Errorf("expected fullName was %s but got %s", c["fullName"], cl.FullName)
		}

		if cl.Id != c["id"] {
			t.Errorf("expected id was %s but got %s", c["id"], cl.Id)
		}
	})
}
