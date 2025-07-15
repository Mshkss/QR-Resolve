package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ResolveHandler(collection *mongo.Collection) echo.HandlerFunc {
	return func(c echo.Context) error {
		mac := c.Param("mac")
		log.Println("ResolveHandler called with mac =", mac)  // <-- вот здесь

		start := time.Now()

		var result struct {
			ID          string `bson:"_id"`
			Mac         string `bson:"mac"`
			RedirectURL string `bson:"redirect_url"`
		}

		dbStart := time.Now()
		err := collection.FindOne(context.Background(), bson.M{"mac": mac}).Decode(&result)
		dbDuration := time.Since(dbStart)

		if err != nil {
			log.Printf("Mongo FindOne error: %v (duration: %v)", err, dbDuration)
			return c.JSON(http.StatusNotFound, map[string]string{"error": "device not found"})
		}

		totalDuration := time.Since(start)
		log.Printf("Request processed: total %v, DB query %v", totalDuration, dbDuration)

		return c.JSON(http.StatusOK, map[string]string{"redirect_url": result.RedirectURL})
	}
}
