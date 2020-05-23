package router

import (
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

// SetupRouting APIのルーティングを行います
func SetupRouting(e *echo.Echo, client *UserProvider) {
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
			apiItems.GET("", GetItems)
			apiItems.POST("", PostItems)
			apiItems.GET("/:id", GetItem)
			apiItems.PUT("/:id", PutItem)
			apiItems.DELETE("/:id", DeleteItem)
			apiItems.POST("/:id/owners", PostOwners)
			apiItems.PUT("/:id/owners", PutOwners)
			apiItems.POST("/:id/logs", PostLogs)
			apiItems.POST("/:id/comments", PostComments)
			apiItems.POST("/:id/likes", PostLikes)
			apiItems.DELETE("/:id/likes", DeleteLikes)
		}

		apiComments := api.Group("/comments")
		{
			apiComments.GET("", GetComments)
		}

		apiFiles := api.Group("/files")
		{
			apiFiles.POST("", PostFile, middleware.BodyLimit("3MB"))
		}

	}
	e.GET("/api/files/:id", GetFile)
}
