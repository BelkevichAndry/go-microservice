package main

import (
	pb "github.com/BelkevichAndry/go-microservice/shippy-user-service/proto/user"
	"github.com/dgrijalva/jwt-go"
)

var (
	key = []byte("mySeperSecretKey")
)

type CustomClaims struct {
	User *pb.User
	jwt.StandardClaims
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(token string) (*CustomClaims, error) {

	tokenType, err := jwt.ParseWithClaims(string(key), &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *pb.User) (string, error) {
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "shippy.service.user",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	return token.SignedString(key)
}
