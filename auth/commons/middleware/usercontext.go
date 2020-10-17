package middleware

import (
	"auth/internal/models"
	"net/http"

	"github.com/adityak368/swissknife/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// EchoUserContextMiddleware returns a echo middleware for setting user context
func EchoUserContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Get("user").(*jwt.Token)
			claims := token.Claims.(jwt.MapClaims)
			user, ok := claims["info"].(*models.User)
			if !ok {
				return response.NewError(http.StatusBadRequest, "InvalidUser")
			}
			c.Set("user", user)
			return next(c)
		}
	}
}
