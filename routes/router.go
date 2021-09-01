package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sternth/go-punch-time/handler"
	"github.com/sternth/go-punch-time/utils"
)

func NewRouter(port string) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", rootHandler)
	e.GET("/ping", pingHandler)
	e.GET("/tasks", handler.GetTasks)
	e.POST("/tasks", handler.CreateTask)
	e.GET("/tasks/:id", handler.GetTask)
	e.PUT("/tasks/:id", handler.UpdateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}

func rootHandler(c echo.Context) error {
	tag, err := utils.GetLatestTag()
	if err != nil {
		return c.String(http.StatusOK, "go-punch-time@<error>")
	}
	return c.String(http.StatusOK, fmt.Sprintf("go-punch-time@%v", tag))
}

func pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
