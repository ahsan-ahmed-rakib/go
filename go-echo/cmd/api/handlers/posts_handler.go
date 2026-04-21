package handlers

import (
	"net/http"
	"strconv"

	"github.com/ahsan-ahmed-rakib/go-echo/cmd/api/service"
	"github.com/labstack/echo"
)

func PostHandler(e echo.Context) error{
	data, err := service.GetAll()
	if err != nil {
		e.String(http.StatusBadGateway, "Unable to process data")
	}
	
	res := make(map[string]any)
	res["ok"] = "ok"
	res["data"] = data
	
	return e.JSON(http.StatusOK, res)
}

func SinglePostHandler(e echo.Context) error {
	id := e.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		e.String(http.StatusBadGateway, "Unable to process data")
	}
	
	data,err := service.GetById(idx)
	if err != nil {
		e.String(http.StatusBadGateway, "Unable to process data")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return e.JSON(http.StatusOK, res)
}