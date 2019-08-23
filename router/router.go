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
		apiUsers := api.Group("/users")
		{
			apiUsers.GET("", GetUsers)
			apiUsers.GET("/me", GetUsersMe)
			apiUsers.PUT("", PutUsers)
		}

		apiItems := api.Group("/items")
		{
			apiItems.POST("", PostItems)
		}
	}
}
