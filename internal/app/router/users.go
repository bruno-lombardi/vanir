package router

import (
	"vanir/internal/app/presentation/adapters"
	controller "vanir/internal/app/presentation/controller/users"
	"vanir/internal/app/presentation/middlewares"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(r *echo.Group) {
	userService := services.GetUserService(repositories.GetUserRepository(), crypto.GetHasher())
	getUserController := controller.NewGetUserController(userService)
	updateUserController := controller.NewUpdateUserController(userService)
	createUserController := controller.NewCreateUserController(userService)
	authenticatedMiddleware := middlewares.GetAuthenticatedMiddleware(crypto.GetEncrypter(), userService)

	r.POST("", adapters.AdaptControllerToEchoJSON(
		createUserController, &models.CreateUserParams{},
	))
	r.PUT("/:id",
		adapters.AdaptControllerToEchoJSON(updateUserController, &models.UpdateUserParams{}),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)
	r.GET("/:id",
		adapters.AdaptControllerToEchoJSON(getUserController, nil),
		adapters.AdaptMiddlewareToEcho(authenticatedMiddleware, nil),
	)
}
