package middlewares

import (
	"net/http"
	"strings"
	"sync"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type AuthenticatedMiddleware struct {
	encrypter   *crypto.Encrypter
	userService *services.UserService
}

var authenticatedMiddleware *AuthenticatedMiddleware
var authenticatedMiddlewareOnce sync.Once

func GetAuthenticatedMiddleware(encrypter crypto.Encrypter, userService services.UserService) *AuthenticatedMiddleware {
	authenticatedMiddlewareOnce.Do(func() {
		authenticatedMiddleware = NewAuthenticatedMiddleware(encrypter, userService)
	})
	return authenticatedMiddleware
}

func NewAuthenticatedMiddleware(encrypter crypto.Encrypter, userService services.UserService) *AuthenticatedMiddleware {
	return &AuthenticatedMiddleware{
		encrypter:   &encrypter,
		userService: &userService,
	}
}

func (m *AuthenticatedMiddleware) Handle(req *protocols.HttpRequest) error {
	authorization := req.HttpReq.Header.Get("authorization")
	if authorization == "" {
		return unauthorized("Invalid credentials provided")
	}

	splitted := strings.Split(authorization, " ")
	token := splitted[1]

	valid, subject := (*m.encrypter).ValidateToken(token)
	if !valid {
		return unauthorized("Invalid credentials provided")
	}

	user, err := (*m.userService).Get(subject)
	if err != nil {
		return unauthorized("Invalid credentials provided")
	}
	req.AuthenticatedUser = user
	return nil
}

func unauthorized(message string) *protocols.AppError {
	return protocols.NewAppError(message, http.StatusUnauthorized)
}
