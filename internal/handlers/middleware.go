package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret []byte) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Println("JWTMiddleware called") // 👈

			// 1. Чтение заголовка Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
			}

			// 2. Удаление "Bearer "
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			// 3. Парсинг и валидация токена
			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				// 3.1 Проверка алгоритма
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return secret, nil
			})
			if err != nil || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}

			// 4. Извлечение claims (например, username)
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
			}
			log.Printf("Claims: %+v", claims)

			// 5. Сохраняем в контекст
			//c.Set("username", claims["username"])

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
			return next(c)
		}
	}
}
