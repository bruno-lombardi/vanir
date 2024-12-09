package controller

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

// CreateUserController godoc
// @Summary Create new user
// @Description Creates a new user with the provided information.
// @Tags users
// @Accept application/json
// @Produce json
// @Param request body models.CreateUserParams true "create user params"
// @Success 200 {object} models.User
// @Router /v1/users [POST]
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
