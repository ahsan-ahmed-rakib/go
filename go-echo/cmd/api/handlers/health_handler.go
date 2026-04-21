package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HealthCheckHandler(e echo.Context) error {
	return e.String(http.StatusOK, "OK!")
}