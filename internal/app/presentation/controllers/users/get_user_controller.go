package controllers

import (
	"net/http"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type GetUserController struct {
	userService *services.UserService
}

func NewGetUserController(userService services.UserService) *CreateUserController {
	return &CreateUserController{
		userService: &userService,
	}
}

func (c *GetUserController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	id := req.PathParams["id"]

	var err error
	user, err := (*c.userService).Get(id)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       user,
	}

	return response, nil
}
