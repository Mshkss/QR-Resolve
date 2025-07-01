package main

import (
	"log"

	"QR_Resolve/api"
	"QR_Resolve/internal/db"
	"QR_Resolve/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	client, err := db.Connect() // используйте свой метод подключения
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}
	defer client.Disconnect(nil)

	apiCollection := client.Database("1").Collection("api")
	usersCollection := client.Database("1").Collection("users")

	server := handlers.NewServer(apiCollection, usersCollection)

	e := echo.New()
	api.RegisterHandlers(e, &server)
	log.Fatal(e.Start("0.0.0.0:8080"))
}
