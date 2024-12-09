package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	_ "vanir/internal/app/docs"
	"vanir/internal/app/router"
	"vanir/internal/pkg/config"
	"vanir/internal/pkg/data/db"
	"vanir/internal/pkg/helpers"

	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title Vanir
// @version 1.0.1
// @description This is an API used to track crypto currencies prices and manage your favorites

// @contact.name Bruno Lombardi
// @contact.url https://github.com/bruno-lombardi

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3333
// @BasePath /api/v1
func Run() {
	config.Setup()
	db.SetupDB()
	conf := config.GetConfig()
	e := echo.New()
	e.Validator = helpers.NewCustomValidator()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, time=${time_rfc3339_nano} method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		fmt.Println(c.Path(), c.QueryParams(), err)
		e.DefaultHTTPErrorHandler(err, c)
	}
	apiV1 := e.Group("/v1")
	router.SetupRootRouter(apiV1)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	go func() {
		if err := e.Start(":" + conf.Server.Port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	e.Logger.Infof("gracefully shutting down")
	close(quit)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
