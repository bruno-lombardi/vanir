package adapters

import (
	"net/http"
	"strings"
	"vanir/internal/pkg/data/models"
	"vanir/internal/pkg/protocols"

	"github.com/labstack/echo/v4"
)

const (
	AUTHENTICATED_USER_KEY = "authenticated_user"
)

func AdaptControllerToEchoJSON(controller protocols.Controller, body interface{}) func(c echo.Context) (err error) {
	return func(c echo.Context) (err error) {
		httpRequest := &protocols.HttpRequest{
			Body: body,
		}
		user, ok := c.Get(AUTHENTICATED_USER_KEY).(*models.User)
		if ok && user != nil {
			httpRequest.AuthenticatedUser = user
		}

		httpRequest.HttpReq = c.Request()

		params := map[string]string{}
		for _, p := range c.ParamNames() {
			for _, v := range c.ParamValues() {
				params[p] = v
			}
		}
		httpRequest.PathParams = params
		httpRequest.QueryParams = c.QueryParams()

		if body != nil {
			if err = c.Bind(httpRequest.Body); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
			if err = c.Validate(httpRequest.Body); err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
			}
		}

		response, err := controller.Handle(httpRequest)
		if err != nil {
			switch err := err.(type) {
			case *protocols.AppError:
				return c.JSON(err.StatusCode, map[string]string{"message": err.Error()})
			default:
				return c.JSON(response.StatusCode, map[string]string{"message": err.Error()})
			}
		}

		for k, v := range response.Headers {
			c.Response().Header().Set(k, strings.Join(v[:], ";"))
		}
		return c.JSON(response.StatusCode, response.Body)
	}
}

func AdaptMiddlewareToEcho(middleware protocols.Middleware, body interface{}) func(echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			httpRequest := &protocols.HttpRequest{
				Body: body,
			}
			httpRequest.HttpReq = c.Request()

			params := map[string]string{}
			for _, p := range c.ParamNames() {
				for _, v := range c.ParamValues() {
					params[p] = v
				}
			}
			httpRequest.PathParams = params
			httpRequest.QueryParams = c.QueryParams()

			if body != nil {
				if err = c.Bind(httpRequest.Body); err != nil {
					return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
				}
			}

			err = middleware.Handle(httpRequest)
			if err != nil {
				switch err := err.(type) {
				case *protocols.AppError:
					return c.JSON(err.StatusCode, map[string]string{"message": err.Error()})
				default:
					return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
				}
			}
			if httpRequest.AuthenticatedUser != nil {
				c.Set(AUTHENTICATED_USER_KEY, httpRequest.AuthenticatedUser)
			}

			next(c)
			return nil
		}
	}

}
