package router

import (
	"vanir/internal/app/presentation/adapters"
	controllers "vanir/internal/app/presentation/controllers/auth"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupAuthRoutes(r *echo.Group) {
	r.POST("", adapters.AdaptControllerToEchoJSON(
		controllers.NewAuthController(services.GetAuthService()), &models.AuthCredentialsDTO{},
	))
}
