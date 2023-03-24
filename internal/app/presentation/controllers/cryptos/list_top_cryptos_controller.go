package controllers

import (
	"net/http"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"
	"vanir/internal/pkg/services"
)

type ListTopCryptosController struct {
	cryptoService *services.CryptoService
}

func NewListTopCryptosController(cryptoService services.CryptoService) *ListTopCryptosController {
	return &ListTopCryptosController{
		cryptoService: &cryptoService,
	}
}

func (c *ListTopCryptosController) Handle(req *protocols.HttpRequest) (*protocols.HttpResponse, error) {
	params := req.Body.(*models.ListTopCryptoCurrenciesQueryParams)

	topCryptos, err := (*c.cryptoService).ListTopCryptoCurrencies(params)
	if err != nil {
		return &protocols.HttpResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	response := &protocols.HttpResponse{
		StatusCode: http.StatusOK,
		Body:       topCryptos,
	}

	return response, nil
}
