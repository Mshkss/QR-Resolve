package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Список защищённых маршрутов: method + path
var protectedRoutes = map[string]bool{
	"POST /api/new":    true,
	"PUT /api/:mac":    true,
	"DELETE /api/:mac": true,
}

func ConditionalJWTMiddleware(secret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			routeKey := c.Request().Method + " " + c.Path()

			if protectedRoutes[routeKey] {
				log.Println("[JWT] Проверка токена для:", routeKey)

				authHeader := c.Request().Header.Get("Authorization")
				if authHeader == "" {
					return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
				}

				tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

				token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method")
					}
					return secret, nil
				})
				if err != nil || !token.Valid {
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
				}

				claims, ok := token.Claims.(jwt.MapClaims)
				if !ok {
					return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
				}

				usernameRaw := claims["username"]
				var username string
				switch v := usernameRaw.(type) {
				case string:
					username = v
				case float64:
					username = fmt.Sprintf("%.0f", v)
				default:
					username = fmt.Sprintf("%v", v)
				}

				c.Set("username", username)
			} else {
				log.Println("[JWT] Пропущен:", routeKey)
			}

			return next(c)
		}
	}
}