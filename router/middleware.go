package router

import (
	"github.com/labstack/echo"

	"github.com/traPtitech/booQ/utils"
)

func middlewareAuthUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		client := utils.TraqClient{}
		c, err := client.GetUsersMe(c)
		if err != nil {
			return err
		}
		return next(c)
	}
}
