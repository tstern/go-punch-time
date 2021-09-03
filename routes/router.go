package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sternth/go-punch-time/handler"
	"github.com/sternth/go-punch-time/utils"
)

func NewRouter(port string) {
	e := echo.New()
	db := utils.ConnectDb()
	taskCtrl := handler.NewTaskController(db)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", handler.RootHandler)
	e.GET("/ping", handler.PingHandler)
	e.GET("/tasks", taskCtrl.GetTasks)
	e.POST("/tasks", handler.CreateTask)
	e.GET("/tasks/:id", handler.GetTask)
	e.PUT("/tasks/:id", handler.UpdateTask)
	e.DELETE("/tasks/:id", handler.DeleteTask)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
