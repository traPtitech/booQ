package router

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo/middleware"
	"github.com/traPtitech/booQ/model"

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
			apiUsers.PUT("", PutUsers, MiddlewareAdmin)
		}

		apiItems := api.Group("/items")
		{
			apiItems.GET("", GetItems)
			apiItems.POST("", PostItems, MiddlewareItemSocial(func(c echo.Context) model.Item {
				item := model.Item{}
				body, err := ioutil.ReadAll(c.Request().Body)
				if err != nil {
					return model.Item{}
				}
				c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
				if err = json.Unmarshal(body, &item); err != nil {
					return model.Item{}
				}
				return item
			}))
			apiItems.GET("/:id", GetItem)
			apiItems.PUT("/:id", PutItem)
			apiItems.DELETE("/:id", DeleteItem, MiddlewareAdmin)
			apiItems.POST("/:id/owners", PostOwners, MiddlewareItemSocial(func(c echo.Context) model.Item {
				itemID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return model.Item{}
				}
				item, _ := model.GetItemByID(uint(itemID))
				return item
			}))
			apiItems.PUT("/:id/owners", PutOwners, MiddlewareItemSocial(func(c echo.Context) model.Item {
				itemID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return model.Item{}
				}
				item, _ := model.GetItemByID(uint(itemID))
				return item
			}))
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
