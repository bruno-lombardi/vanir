package router

import (
	"vanir/internal/app/presentation/adapters"
	controllers "vanir/internal/app/presentation/controllers/favorite"
	"vanir/internal/app/presentation/middlewares"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupFavoritesRoutes(r *echo.Group) {
	favoriteService := services.GetFavoriteService()
	authenticatedMiddleware := middlewares.GetAuthenticatedMiddleware()
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
