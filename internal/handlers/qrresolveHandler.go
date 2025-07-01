package handlers

import (
	"QR_Resolve/api"
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

	var entry api.ApiEntry
	err := s.ApiCollection.FindOne(ctx.Request().Context(), filter).Decode(&entry)
	if err != nil {
		log.Printf("MongoDB FindOne error: %v", err)
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "not found"})
	}

	log.Printf("MongoDB found entry: %+v", entry)
	return ctx.JSON(http.StatusOK, api.ResolvedLink{Url: &entry.RedirectUrl})
}
