package myjwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	tokenTTL   = 10 * time.Minute // Token
	signingKey = "kdmckdmkcmdmcjnfjvnjfncmc"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

// generating jwt access token by id (just int value)
func GenerateToken(id int) (string, error) {
	// get user from db
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		id,
	})
	return token.SignedString([]byte(signingKey))
}

// parse jwt access token  and return UserID if it's valid or an error else
func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // signing method is HMAC?
		if !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("cannot cast claim to token claims")
	}
	return claims.UserId, nil
}
