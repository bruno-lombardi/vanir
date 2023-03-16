package controllers

import (
	"net/http"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type UpdateUserController struct {
	userService *services.UserService
}

func NewUpdateUserController(userService services.UserService) *UpdateUserController {
	return &UpdateUserController{
		userService: &userService,
	}
}

func (c *UpdateUserController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	updateUserDTO := req.Body.(*models.UpdateUserDTO)

	var err error
	user, err := (*c.userService).Update(updateUserDTO)
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
