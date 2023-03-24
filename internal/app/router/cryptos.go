package router

import (
	"vanir/internal/app/presentation/adapters"
	controllers "vanir/internal/app/presentation/controllers/cryptos"
	"vanir/internal/app/presentation/middlewares"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupCryptoRoutes(r *echo.Group) {
	cryptoService := services.GetCryptoService()
	authenticatedMiddleware := middlewares.GetAuthenticatedMiddleware()
	listTopCryptosController := controllers.NewListTopCryptosController(cryptoService)
	r.GET("",
		adapters.AdaptControllerToEchoJSON(listTopCryptosController, &models.ListTopCryptoCurrenciesQueryParams{}),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)
}
