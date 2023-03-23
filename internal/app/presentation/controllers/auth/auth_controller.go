package controllers

import (
	"net/http"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: &authService,
	}
}

func (c *AuthController) Handle(req *protocols.HttpRequest) (res *protocols.HttpResponse, err error) {
	authCredentialsDTO := req.Body.(*models.AuthCredentialsDTO)

	authenticationResponse, err := (*c.authService).Authenticate(authCredentialsDTO)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       authenticationResponse,
	}

	return response, nil
}
