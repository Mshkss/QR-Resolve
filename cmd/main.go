package main

import (
	"log"

	"qr-resolve/internal/db"
	"qr-resolve/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Подключение к Mongo
	client, collection := db.Connect()
	defer client.Disconnect(nil)

	// Echo instance
	e := echo.New()

	// Роут
	e.GET("/resolve/:mac", handlers.ResolveHandler(collection))

	// Запуск
	log.Println("Сервер запущен на :8080")

	e.Logger.Fatal(e.Start(":8080"))
}
