package middleware

import (
	"github.com/labstack/echo"
)

// ServerHeader function
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "apigateway/2.0 (AArch64)")
		return next(c)
	}
}
