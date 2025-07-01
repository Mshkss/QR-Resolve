package handlers

import (
	"QR_Resolve/api"
	"net/http"

	"github.com/labstack/echo/v4"
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

// Реализация методов интерфейса:

func (s *Server) AddApiEntry(ctx echo.Context) error {
	// TODO: реализовать добавление записи
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

func (s *Server) DeleteApiEntry(ctx echo.Context, mac string) error {
	// TODO: реализовать удаление записи
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

func (s *Server) UpdateApiEntry(ctx echo.Context, mac string) error {
	// TODO: реализовать обновление записи
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

func (s *Server) LoginUser(ctx echo.Context) error {
	// TODO: реализовать авторизацию
	return ctx.JSON(http.StatusNotImplemented, map[string]string{"error": "not implemented"})
}

// func (s *Server) GetPing(ctx echo.Context) error {
// 	return ctx.JSON(http.StatusOK, api.Pong{Ping: "pong"})
// }

