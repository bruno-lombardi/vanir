package router

import (
	"vanir/internal/app/presentation/adapters"
	controllers "vanir/internal/app/presentation/controllers/cryptos"
	"vanir/internal/app/presentation/middlewares"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/http/clients"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupCryptoRoutes(r *echo.Group) {
	userService := services.GetUserService(repositories.GetUserRepository(), crypto.GetHasher())
	cryptoService := services.GetCryptoService(clients.NewCryptoCompareHttpClient(), repositories.GetFavoritesRepository())
	authenticatedMiddleware := middlewares.GetAuthenticatedMiddleware(crypto.GetEncrypter(), userService)
	listTopCryptosController := controllers.NewListTopCryptosController(cryptoService)
	favoriteCryptosController := controllers.NewListUserFavoriteCryptosController(cryptoService)

	r.GET("/toplist",
		adapters.AdaptControllerToEchoJSON(listTopCryptosController, &models.ListTopCryptoCurrenciesQueryParams{}),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)
	r.GET("/favorites",
		adapters.AdaptControllerToEchoJSON(favoriteCryptosController, &models.ListUserFavoriteCryptoCurrenciesQueryParams{}),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)
}
