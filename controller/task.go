package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sternth/go-punch-time/model"
	"github.com/sternth/go-punch-time/store"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	query := e.Request().URL.Query()
	lastDays, err := strconv.Atoi(query.Get("lastDays"))
	if err != nil {
		lastDays = 100
	}

	ctx := e.Request().Context()
	tasks, err = ctrl.store.FindAll(ctx, lastDays)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, tasks)
}

func (ctrl *TaskController) Create(e echo.Context) error {
	var task model.Task

	if err := e.Bind(&task); err != nil {
		return err
	}

	ctx := e.Request().Context()
	if err := ctrl.store.Create(ctx, &task); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusCreated, task)
}

func (ctrl *TaskController) GetTask(e echo.Context) error {
	var task *model.Task

	id := e.Param("id")
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e.String(http.StatusBadRequest, fmt.Sprintf("document id %s is invalid", id))
	}

	ctx := e.Request().Context()
	task, err = ctrl.store.Find(ctx, docId)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return e.String(http.StatusNotFound, fmt.Sprintf("task with id %s not found", id))
		} else {
			return e.String(http.StatusInternalServerError, err.Error())
		}
	}

	return e.JSON(http.StatusOK, *task)
}

func (ctrl *TaskController) Update(e echo.Context) error {
	var task *model.Task

	id := e.Param("id")
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e.String(http.StatusBadRequest, fmt.Sprintf("document id %s is invalid", id))
	}

	if err := e.Bind(&task); err != nil {
		return err
	}

	ctx := e.Request().Context()
	if err = ctrl.store.Update(ctx, docId, task); err != nil {
		if err == mongo.ErrNoDocuments {
			return e.String(http.StatusNotFound, fmt.Sprintf("task with id %s not found", id))
		} else {
			return e.String(http.StatusInternalServerError, err.Error())
		}
	}

	return e.JSON(http.StatusOK, task)
}

func (ctrl *TaskController) Delete(e echo.Context) error {
	id := e.Param("id")
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e.String(http.StatusBadRequest, fmt.Sprintf("document id %s is invalid", id))
	}

	ctx := e.Request().Context()
	if err = ctrl.store.Delete(ctx, docId); err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.String(http.StatusOK, fmt.Sprintf("task with id %s deleted", id))
}
