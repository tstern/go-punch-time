package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sternth/go-punch-time/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Start int64              `json:"start"`
	End   int64              `json:"end"`
	Text  string             `json:"text"`
	Type  string             `json:"type"`
}

func GetTasks(e echo.Context) error {
	ctx := e.Request().Context()
	db := utils.ConnectDb()
	collection := db.Collection("tasks")
	lastDate := getLastDate(e)
	filter := bson.M{"start": bson.M{"$gte": lastDate}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var tasks []Task
	if err = cursor.All(ctx, &tasks); err != nil {
		log.Fatal(err)
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
