package controllers

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte("THEISPRIVATE")

func CreateToken(fname string, lname string, email string) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"FirstName": fname,
		"LastName":  lname,
		"Email":     email,
		"iat":       time.Now().Unix(),
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := claims.SignedString(secret_key)
	if err != nil {
		return ""
	}
	return tokenString
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret_key, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
