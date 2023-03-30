package controllers

import (
	"net/http"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type RemoveFavoriteController struct {
	favoriteService *services.FavoriteService
}

func NewRemoveFavoriteController(favoriteService services.FavoriteService) *RemoveFavoriteController {
	return &RemoveFavoriteController{
		favoriteService: &favoriteService,
	}
}

func (c *RemoveFavoriteController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	params := &models.DeleteFavoriteParams{}
	params.Reference = req.PathParams["reference"]
	params.UserID = req.AuthenticatedUser.ID

	err := (*c.favoriteService).Delete(params)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusNoContent,
		Body:       nil,
	}

	return response, nil
}
