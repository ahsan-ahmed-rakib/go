package main

import (
	"github.com/ahsan-ahmed-rakib/go-echo/cmd/api/handlers"
	"github.com/labstack/echo"
)

func main () {
	e := echo.New()
	e.GET("health-check", handlers.HealthCheckHandler)
	e.GET("posts", handlers.PostHandler)
	e.GET("post/:id", handlers.SinglePostHandler)

	e.Logger.Fatal(e.Start(":5000"))
}
