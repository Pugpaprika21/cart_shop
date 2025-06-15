package jwtx

import (
	"miniservice/app/internal/enum"
	"miniservice/app/pkg/response"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type IJwtx interface {
	Validate() echo.MiddlewareFunc
}

type jwtx struct {
	secret string
}

func New() IJwtx {
	return &jwtx{secret: os.Getenv("AUTH_TOKEN")}
}

func (j *jwtx) Validate() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var resp = response.NewResponseBuilder()

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, resp.Code(enum.FOR_AUTH_ERROR).Message("Missing Authorization header").Build())
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader {
				return c.JSON(http.StatusUnauthorized, resp.Code(enum.FOR_AUTH_ERROR).Message("Invalid token format").Build())
			}

			if tokenString != j.secret {
				return c.JSON(http.StatusUnauthorized, resp.Code(enum.FOR_AUTH_ERROR).Message("Invalid token").Build())
			}

			c.Set("token", j.secret)
			return next(c)
		}
	}
}
