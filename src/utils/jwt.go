package utils

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var key = []byte("abcdefg")

var claims = &jwt.StandardClaims{
	ExpiresAt: time.Now().Unix() + 24*3600,
	Issuer:    "lmm",
	Id:        "test",
}

func GenerateToken() string {
	Ptoken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)

	token, err := Ptoken.SignedString(key)
	if err != nil {
		log.Panic(err)
	}
	return token
}

func VerifyToken(token string) *jwt.Token {
	Ptoken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	// if claims, ok := Ptoken.Claims.(jwt.StandardClaims); ok && Ptoken.Valid {
	// 	return claims.Issuer
	// }

	return Ptoken
	// Ptoken.Claims
}
