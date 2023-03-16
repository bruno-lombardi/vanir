package router

import (
	"vanir/internal/app/presentation/adapters"
	"vanir/internal/app/presentation/controllers"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(r *echo.Group) {
	r.POST("/users", adapters.AdaptEchoJSON(
		controllers.NewCreateUserController(*services.GetUserService()), &models.CreateUserDTO{},
	))
	r.PUT("/users/:id", adapters.AdaptEchoJSON(
		controllers.NewUpdateUserController(*services.GetUserService()), &models.UpdateUserDTO{},
	))
}
