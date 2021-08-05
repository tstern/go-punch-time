package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(port string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello darkness my old friend.")
	})
	e.GET("/port", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Running on port %v", port))
	})
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	e.Any("/any", func(c echo.Context) error {
		return c.String(http.StatusOK, "any...")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
