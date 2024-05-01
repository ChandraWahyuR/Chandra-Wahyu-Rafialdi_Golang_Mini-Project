package authorization

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func OnlyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Token from jwt get role user
		token := c.Get("user").(*jwt.Token)

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}
		// Role authorization for admin only
		userType, ok := claims["role"].(string)
		if !ok || userType != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "Eror not alowed, admin only")
		}

		return next(c)
	}
}
