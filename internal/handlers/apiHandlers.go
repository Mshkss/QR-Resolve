package handlers

import (
	"QR_Resolve/api"
	"QR_Resolve/internal/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
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
	// todo: НЕ РАБОТАЕТ JWT >>>
	// user := ctx.Get("username")
	// log.Printf("JWT user: %v", user)
	// username, ok := user.(string)
	// if !ok || username == "" {
	// 	return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid JWT claims"})
	// }

	// log.Printf("JWT username: %s", username)
	// <<<

	var req models.ApiEntry
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// 2. Проверяем, существует ли запись с таким же MAC-адресом
	var existingDevice models.ApiEntry
	err := s.ApiCollection.FindOne(ctx.Request().Context(), bson.M{
		"mac":          req.Mac,
		"redirect_url": req.RedirectUrl,
	}).Decode(&existingDevice)
	if err == nil {
		return ctx.JSON(http.StatusConflict, map[string]string{"error": "entry with this MAC and URL already exists"})
	}

	// 3. Вставляем новую запись в монгу
	_, err = s.ApiCollection.InsertOne(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to insert entry"})
	}

	return ctx.JSON(http.StatusCreated, map[string]string{"result": "Mac added"})
}

// DELETE /api/:mac"
func (s *Server) DeleteApiEntry(ctx echo.Context, mac string) error {
	// TODO: проверка JWT
	var req models.ApiEntry
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	s.ApiCollection.DeleteMany(ctx.Request().Context(), bson.M{"mac": mac})

	return ctx.JSON(http.StatusCreated, map[string]string{"result": "All Mac deleted"})

}

// PUT /api/:mac
func (s *Server) UpdateApiEntry(ctx echo.Context, mac string) error {
	// TODO: проверка JWT
	var req models.ApiEntry
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		}
	
	_, err := s.ApiCollection.UpdateOne(
		ctx.Request().Context(),
		bson.M{"mac": mac},
		bson.M{"$set": bson.M{
			"redirect_url": req.RedirectUrl,
		}},
	)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update entry"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"result": "Mac updated"})
}
