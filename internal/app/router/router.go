package router

import "github.com/labstack/echo/v4"

func SetupRootRouter(r *echo.Group) {
	SetupUserRoutes(r.Group("/users"))
	SetupAuthRoutes(r.Group("/auth"))
}
