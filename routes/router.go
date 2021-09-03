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
	e.GET("/tasks", taskCtrl.GetAll)
	e.POST("/tasks", taskCtrl.Create)
	e.GET("/tasks/:id", taskCtrl.GetTask)
	e.PUT("/tasks/:id", taskCtrl.Update)
	e.DELETE("/tasks/:id", taskCtrl.Delete)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", port)))
}
