package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sternth/go-punch-time/model"
	"github.com/sternth/go-punch-time/store"
	"github.com/sternth/go-punch-time/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskController struct {
	store *store.TaskStore
}

func NewTaskController(db *mongo.Database) *TaskController {
	collection := db.Collection("tasks")
	return &TaskController{
		store: store.NewTaskStore(collection),
	}
}

func (ctrl *TaskController) GetAll(e echo.Context) error {
	var tasks []model.Task

	ctx := e.Request().Context()
	query := e.Request().URL.Query()

	lastDays, err := strconv.Atoi(query.Get("lastDays"))
	if err != nil {
		lastDays = 100
	}

	tasks, err = ctrl.store.FindAll(ctx, lastDays)
	if err != nil {
		e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, tasks)
}

func (ctrl *TaskController) Create(e echo.Context) error {
	db := utils.ConnectDb()
	return e.String(http.StatusOK, fmt.Sprintf("%v - CREATE task", db.Name()))
}

func (ctrl *TaskController) GetTask(e echo.Context) error {
	db := utils.ConnectDb()
	id := e.Param("id")
	return e.String(http.StatusOK, fmt.Sprintf("%v - GET task id: %v", db.Name(), id))
}

func (ctrl *TaskController) Update(e echo.Context) error {
	db := utils.ConnectDb()
	id := e.Param("id")
	return e.String(http.StatusOK, fmt.Sprintf("%v - UPDATE task id: %v", db.Name(), id))
}

func (ctrl *TaskController) Delete(e echo.Context) error {
	db := utils.ConnectDb()
	id := e.Param("id")
	return e.String(http.StatusOK, fmt.Sprintf("%v - DELETE task id: %v", db.Name(), id))
}
