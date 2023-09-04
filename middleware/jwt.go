package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go-rest-api-server/domain"
	"go-rest-api-server/service"
	"net/http"
	"strings"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the jwt token from the header
		authHeader := c.Request().Header.Get("Authorization")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}
		tokenString := parts[1]

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &domain.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Make sure the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return service.SecretKey, nil
		})

		// Check if there was an error in parsing...
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		// If the token is valid, save the claims in context
		if claims, ok := token.Claims.(*domain.JWTClaims); ok && token.Valid {
			c.Set("id", claims.ID)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
		}

		// Call the next handler
		return next(c)
	}
}
