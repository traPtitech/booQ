package router

import (
	"net/http"

	"github.com/labstack/echo"
)

// SetupRouting APIのルーティングを行います
func SetupRouting(e *echo.Echo, client Traq) {
	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	api := e.Group("/api", client.MiddlewareAuthUser)
	{
		apiUser := api.Group("/user")
		{
			apiUser.GET("/me", GetUserMe)
		},

		apiUsers := api.Group("/users")
		{
			apiUsers.PUT("/", PutUser)
		}
	}
}
