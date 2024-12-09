package controllers

import (
	"net/http"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type ListUserFavoriteCryptosController struct {
	cryptoService *services.CryptoService
}

func NewListUserFavoriteCryptosController(cryptoService services.CryptoService) *ListUserFavoriteCryptosController {
	return &ListUserFavoriteCryptosController{
		cryptoService: &cryptoService,
	}
}

func (c *ListUserFavoriteCryptosController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	params := req.Body.(*models.ListUserFavoriteCryptoCurrenciesQueryParams)
	params.UserID = req.AuthenticatedUser.ID

	favoriteCryptosResponse, err := (*c.cryptoService).ListUserFavoriteCryptoCurrencies(params)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       favoriteCryptosResponse,
	}

	return response, nil
}
