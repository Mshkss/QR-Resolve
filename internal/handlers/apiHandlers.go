package handlers

import (
	"QR_Resolve/api"
	"QR_Resolve/internal/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get /resolve/{category}/{deviceId}
func (s *Server) ResolveDevice(ctx echo.Context, category string, deviceId string) error {
	log.Printf("ResolveDevice called: category=%s, deviceId=%s", category, deviceId)

	filter := map[string]interface{}{
		"mac": deviceId,
		// "category": category,
	}
	log.Printf("MongoDB filter: %+v", filter)

	var entry models.ApiEntry
	err := s.ApiCollection.FindOne(ctx.Request().Context(), filter).Decode(&entry)
	if err != nil {
		log.Printf("MongoDB FindOne error: %v", err)
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}

	log.Printf("MongoDB found entry: %+v", entry)
	log.Printf("MongoDB found entry2: %+v", entry.RedirectUrl)

	return ctx.JSON(http.StatusOK, api.ResolvedLink{Url: &entry.RedirectUrl})

}

// POST /api/new
func (s *Server) AddApiEntry(ctx echo.Context) error {
	user := ctx.Get("username")
	log.Printf("JWT user: %v", user)
	username, ok := user.(string)
	if !ok || username == "" {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid JWT claims"})
	}

	log.Printf("JWT username: %s", username)

	// 2. Парсим тело запроса
	var req models.ApiEntry
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// 3. Вставляем запись в MongoDB
	_, err := s.ApiCollection.InsertOne(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to insert entry"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"result": "entry added"})
}

// 4
// 3
// 1
// 2
// 1
// TODO:
// DELETE /api/:mac"
func (s *Server) DeleteApiEntry(ctx echo.Context, mac string) error {
	// TODO: реализовать удаление записи
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

// PUT /api/:mac
func (s *Server) UpdateApiEntry(ctx echo.Context, mac string) error {
	// TODO: реализовать обновление записи
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}
