package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Avatar   string `json:"avatar"`
	FullName string `json:"fullName"`
	jwt.StandardClaims
}

func GenerateJWTToken(id, avatar, fullName string) string {
	c := Claims{
		Avatar:   avatar,
		FullName: fullName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).UnixNano(),
			Id:        id,
			IssuedAt:  time.Now().UnixNano(),
			Issuer:    "fake_store_auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	ss, _ := token.SignedString([]byte("FAKE_STORE_AUTH"))

	return ss
}
