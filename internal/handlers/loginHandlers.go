package handlers

import (
	"QR_Resolve/api"
	"QR_Resolve/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) LoginUser(ctx echo.Context) error {
	var req api.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	log.Printf("api.LoginRequest=%s", &req)

	// Поиск пользователя по username
	var user models.LoginRequest
	log.Printf("models.LoginRequest=%d", &user.Username) // QR-ResolveV2  | 2025/07/02 12:46:16 models.LoginRequest=274880119504 - что это ? //TODO:

	filter := bson.M{"username": req.Username}
	err := s.UsersCollection.FindOne(ctx.Request().Context(), filter).Decode(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
	}

	// Проверка пароля hash
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password")
	// }

	// проверка пароля !hash
	if user.Password != req.Password {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password")
	}

	secret := []byte("your-secret-key")
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Could not generate token")
	}
	return ctx.JSON(http.StatusOK, api.LoginResponse{Token: &tokenString})
}

// регистрация шифрование:
// hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
