package handlers

import (
	"QR_Resolve/api"
	//"net/http"

	//"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	ApiCollection   *mongo.Collection
	UsersCollection *mongo.Collection
}

func NewServer(apiCol, usersCol *mongo.Collection) Server {
	return Server{
		ApiCollection:   apiCol,
		UsersCollection: usersCol,
	}
}
