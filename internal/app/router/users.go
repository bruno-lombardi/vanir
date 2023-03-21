package router

import (
	"vanir/internal/app/presentation/adapters"
	"vanir/internal/app/presentation/controllers/users"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(r *echo.Group) {
	r.POST("/", adapters.AdaptControllerToEchoJSON(
		controllers.NewCreateUserController(services.GetUserService()), &models.CreateUserDTO{},
	))
	r.PUT("/:id", adapters.AdaptControllerToEchoJSON(
		controllers.NewUpdateUserController(services.GetUserService()), &models.UpdateUserDTO{},
	))
}
