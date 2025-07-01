package handlers

import (
	"QR_Resolve/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GET /ping
func (Server) GetPing(ctx echo.Context) error {
	resp := api.Pong{
		Ping: "pong",
	}

	return ctx.JSON(http.StatusOK, resp)
}
