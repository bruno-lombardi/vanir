package controllers

import (
	"net/http"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type AddFavoriteController struct {
	favoriteService *services.FavoriteService
}

func NewAddFavoriteController(favoriteService services.FavoriteService) *AddFavoriteController {
	return &AddFavoriteController{
		favoriteService: &favoriteService,
	}
}

func (c *AddFavoriteController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	params := req.Body.(*models.CreateFavoriteParams)
	params.UserID = req.AuthenticatedUser.ID

	favorite, err := (*c.favoriteService).Create(params)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       favorite,
	}

	return response, nil
}
