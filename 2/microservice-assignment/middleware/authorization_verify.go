package middleware

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
)

// BearerVerify function to verify token
func BearerVerify() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			header := req.Header
			auth := header.Get("Authorization")
			if auth != os.Getenv("BASIC_AUTH") {
				return echo.NewHTTPError(http.StatusUnauthorized, "authorization is invalid")
			}
			return next(c)
		}
	}
}
