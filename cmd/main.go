package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/scallyt/preact/handler"
	"github.com/scallyt/preact/utils"
)


func main() {
	e := echo.New()

  e.Use(middleware.Logger())
  e.Renderer = utils.NewRenderer("./view/*.html", true)

  //! ROUTES
  e.GET("/", handler.Test)

  e.Logger.Fatal(e.Start(":8080"))
}
