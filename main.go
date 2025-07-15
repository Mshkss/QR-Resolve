package main

import (
	"log"

	"QR_Resolve/api"
	"QR_Resolve/internal/db"
	"QR_Resolve/internal/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	client, err := db.Connect()
	if err != nil {
		log.Fatal("Mongo connection error:", err)
	}
	defer client.Disconnect(nil)

	apiCollection := client.Database("1").Collection("api")
	usersCollection := client.Database("1").Collection("users")

	server := handlers.NewServer(apiCollection, usersCollection)

	e := echo.New()

	e.Use(handlers.ConditionalJWTMiddleware([]byte("your-secret-key")))
	api.RegisterHandlersWithBaseURL(e, &server, "")

	// TODO: весь код ниже переписать под использование oapi codegen net/http middleware
	// >>
	//v2:
	// apiGroup := e.Group("/api", handlers.JWTMiddleware([]byte("your-secret-key")))
	// api.RegisterHandlersWithBaseURL(apiGroup, &server, "/api")
	// api.RegisterHandlersWithBaseURL(e, &server, "/resolve")
	// <<

	//v1:
	//api.RegisterHandlers(e, &server)

	log.Fatal(e.Start("0.0.0.0:8080"))
}
