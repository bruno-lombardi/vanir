package controllers

import (
	"net/http"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type CreateUserController struct {
	userService *services.UserService
}

func NewCreateUserController(userService services.UserService) *CreateUserController {
	return &CreateUserController{
		userService: &userService,
	}
}

func (c *CreateUserController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	createUserParams := req.Body.(*models.CreateUserParams)

	var err error
	user, err := (*c.userService).Create(createUserParams)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusCreated,
		Body:       user,
	}

	return response, nil
}
