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

//
//
//
//
//
//
//
//
//
//TODO:

// POST /api
func (s *Server) AddApiEntry(ctx echo.Context) error {
	// TODO: реализовать добавление записи
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

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
