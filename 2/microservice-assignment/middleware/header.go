package middleware

import (
	"github.com/labstack/echo"
	"os"
)

// ServerHeader function
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, os.Getenv("SERVICE_NAME"))
		return next(c)
	}
}
