package router

import (
	"vanir/internal/app/presentation/adapters"
	controllers "vanir/internal/app/presentation/controller/favorite"
	"vanir/internal/app/presentation/middlewares"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupFavoritesRoutes(r *echo.Group) {
	userService := services.GetUserService(repositories.GetUserRepository(), crypto.GetHasher())
	favoriteService := services.GetFavoriteService(repositories.GetFavoritesRepository())
	authenticatedMiddleware := middlewares.GetAuthenticatedMiddleware(crypto.GetEncrypter(), userService)
	addFavoriteController := controllers.NewAddFavoriteController(favoriteService)
	removeFavoriteController := controllers.NewRemoveFavoriteController(favoriteService)

	r.POST("",
		adapters.AdaptControllerToEchoJSON(addFavoriteController, &models.CreateFavoriteParams{}),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)

	r.DELETE("/:reference",
		adapters.AdaptControllerToEchoJSON(removeFavoriteController, nil),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)
}
