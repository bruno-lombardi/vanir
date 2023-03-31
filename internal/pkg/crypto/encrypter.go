package crypto

import (
	"sync"
	"time"
	"vanir/internal/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

type Encrypter interface {
	CreateToken(subject string) (token string, err error)
	ValidateToken(tokenString string) (valid bool, subject string)
}
type StringSigner interface {
	SignedString(key interface{}) (string, error)
}

type JWTEncrypter struct{}

var jwtEncrypter *JWTEncrypter
var encrypterOnce sync.Once
var newSigner func(method jwt.SigningMethod, claims jwt.Claims) StringSigner
var keyFunc jwt.Keyfunc
var parseToken func(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (*jwt.Token, error)

func GetEncrypter() Encrypter {
	encrypterOnce.Do(func() {
		jwtEncrypter = &JWTEncrypter{}
		newSigner = func(method jwt.SigningMethod, claims jwt.Claims) StringSigner {
			return jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		}
		parseToken = jwt.Parse
		keyFunc = func(t *jwt.Token) (interface{}, error) {
			conf := config.GetConfig()
			return []byte(conf.Server.Secret), nil
		}
	})
	return jwtEncrypter
}

func (e *JWTEncrypter) CreateToken(subject string) (token string, err error) {
	conf := config.GetConfig()
	var signer StringSigner

	claims := jwt.RegisteredClaims{}
	claims.Issuer = "vanir"
	claims.Subject = subject
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 24))
	signer = newSigner(jwt.SigningMethodHS512, claims)
	token, err = signer.SignedString([]byte(conf.Server.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (e *JWTEncrypter) ValidateToken(tokenString string) (valid bool, subject string) {
	token, err := parseToken(tokenString, keyFunc)
	if err != nil {
		return false, ""
	}

	subject, _ = token.Claims.GetSubject()

	return token.Valid, subject
}
