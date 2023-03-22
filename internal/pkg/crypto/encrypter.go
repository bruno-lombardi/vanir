package crypto

import (
	"fmt"
	"time"
	"vanir/internal/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

type Encrypter interface {
	CreateToken(subject string) (token string, err error)
	ValidateToken(tokenString string) (valid bool, subject string)
}

type JWTEncrypter struct{}

var jwtEncrypter *JWTEncrypter

func GetEncrypter() Encrypter {
	if jwtEncrypter == nil {
		jwtEncrypter = &JWTEncrypter{}
	}
	return jwtEncrypter
}

func (e *JWTEncrypter) CreateToken(subject string) (token string, err error) {
	conf := config.GetConfig()

	claims := jwt.RegisteredClaims{}
	claims.Issuer = "vanir"
	claims.Subject = subject
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err = at.SignedString([]byte(conf.Server.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (e *JWTEncrypter) ValidateToken(tokenString string) (valid bool, subject string) {
	conf := config.GetConfig()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(conf.Server.Secret), nil
	})
	if err != nil {
		return false, ""
	}

	subject, err = token.Claims.GetSubject()

	return token.Valid, subject
}
