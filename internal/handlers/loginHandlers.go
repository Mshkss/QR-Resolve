package handlers

import (
	"QR_Resolve/api"
	"QR_Resolve/internal/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *Server) LoginUser(ctx echo.Context) error {
	var req api.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	log.Printf("ResolveDevice called: category=%s", &req)

	// Поиск пользователя по username
	var user models.LoginRequest
	filter := bson.M{"username": req.Username}
	err := s.UsersCollection.FindOne(ctx.Request().Context(), filter).Decode(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
	}

	// Здесь нужно добавить проверить пароль и сгенерировать токен
	token := "some-jwt-token"
	return ctx.JSON(http.StatusOK, api.LoginResponse{Token: &token}) //add some jwt token ??
}
