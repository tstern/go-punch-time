package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
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

func (ctrl *TaskController) GetTasks(e echo.Context) error {
	tasks, err := ctrl.store.FindAll(e.Request())
	if err != nil {
		e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, tasks)
}

func CreateTask(e echo.Context) error {
	db := utils.ConnectDb()
	return e.String(http.StatusOK, fmt.Sprintf("%v - CREATE task", db.Name()))
}

func GetTask(e echo.Context) error {
	db := utils.ConnectDb()
	id := e.Param("id")
	return e.String(http.StatusOK, fmt.Sprintf("%v - GET task id: %v", db.Name(), id))
}

func UpdateTask(e echo.Context) error {
	db := utils.ConnectDb()
	id := e.Param("id")
	return e.String(http.StatusOK, fmt.Sprintf("%v - UPDATE task id: %v", db.Name(), id))
}

func DeleteTask(e echo.Context) error {
	db := utils.ConnectDb()
	id := e.Param("id")
	return e.String(http.StatusOK, fmt.Sprintf("%v - DELETE task id: %v", db.Name(), id))
}

// Get last date by parsing url query parameter "lastDays", by default last date is -100 days from now
func getLastDate(e echo.Context) int64 {
	lastDays := e.Request().URL.Query().Get("lastDays")
	days, err := strconv.Atoi(lastDays)
	if err != nil {
		days = -100
	} else {
		days = -days
	}
	lastUnix := time.Now().AddDate(0, 0, days).Unix()
	return lastUnix * 1000
}
