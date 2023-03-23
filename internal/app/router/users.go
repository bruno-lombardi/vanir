package router

import (
	"vanir/internal/app/presentation/adapters"
	controllers "vanir/internal/app/presentation/controllers/users"
	"vanir/internal/app/presentation/middlewares"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(r *echo.Group) {
	userService := services.GetUserService()
	getUserController := controllers.NewGetUserController(userService)
	updateUserController := controllers.NewUpdateUserController(userService)
	createUserController := controllers.NewCreateUserController(userService)
	authenticatedMiddleware := middlewares.GetAuthenticatedMiddleware()

	r.POST("/", adapters.AdaptControllerToEchoJSON(
		createUserController, &models.CreateUserDTO{},
	))
	r.PUT("/:id",
		adapters.AdaptControllerToEchoJSON(updateUserController, &models.UpdateUserDTO{}),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)
	r.GET("/:id",
		adapters.AdaptControllerToEchoJSON(getUserController, nil),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)
}
