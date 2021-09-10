package route

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sternth/go-punch-time/controller"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewRouter(port string, db *mongo.Database) error {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	commonCtrl := controller.NewCommonController()
	e.GET("/", commonCtrl.Root)
	e.GET("/ping", commonCtrl.Pong)

	taskCtrl := controller.NewTaskController(db)
	e.GET("/tasks", taskCtrl.GetAll)
	e.POST("/tasks", taskCtrl.Create)
	e.GET("/tasks/:id", taskCtrl.GetTask)
	e.PUT("/tasks/:id", taskCtrl.Update)
	e.DELETE("/tasks/:id", taskCtrl.Delete)

	return e.Start(fmt.Sprintf(":%v", port))
}
