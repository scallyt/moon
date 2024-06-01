package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/scallyt/preact/types"
)

func Test(c echo.Context) error {
  data := types.M{"name": "alma"}
  return c.Render(http.StatusOK, "index.html",data)
}
