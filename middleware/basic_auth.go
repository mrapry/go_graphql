package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func BasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "bhinneka" && password == "da1c25d8-37c8-41b1-afe2-42dd4825bfea" {
			return true, nil
		}
		return false, nil
	})
}
