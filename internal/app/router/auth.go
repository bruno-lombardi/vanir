package router

import (
	"vanir/internal/app/presentation/adapters"
	controllers "vanir/internal/app/presentation/controller/auth"
	"vanir/internal/pkg/crypto"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/data/repositories"
	"vanir/internal/pkg/services"

	"github.com/labstack/echo/v4"
)

func SetupAuthRoutes(r *echo.Group) {
	r.POST("", adapters.AdaptControllerToEchoJSON(
		controllers.NewAuthController(
			services.GetAuthService(
				repositories.GetUserRepository(),
				crypto.GetHasher(),
				crypto.GetEncrypter())),
		&models.AuthCredentials{},
	))
}
