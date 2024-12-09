package controller

import (
	"net/http"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type GetUserController struct {
	userService *services.UserService
}

func NewGetUserController(userService services.UserService) *GetUserController {
	return &GetUserController{
		userService: &userService,
	}
}

// GetUserController godoc
// @Summary Get a user by its ID
// @Description Gets an existent user by its ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /v1/users/{id} [GET]
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
