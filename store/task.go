package store

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/sternth/go-punch-time/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskStore struct {
	collection *mongo.Collection
}

func NewTaskStore(collection *mongo.Collection) *TaskStore {
	return &TaskStore{
		collection: collection,
	}
}

func (store *TaskStore) Find(req *http.Request) (*models.Task, error) {
	ctx := req.Context()
	defer ctx.Done()
	return nil, errors.New("Not Implemented")
}

func (store *TaskStore) FindAll(req *http.Request) (*[]models.Task, error) {
	ctx := req.Context()
	lastDate := getLastDate(req.URL.Query())
	filter := bson.M{"start": bson.M{"$gte": lastDate}}
	cursor, err := store.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var tasks []models.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	if tasks == nil {
		tasks = []models.Task{}
	}
	return &tasks, nil
}

func (store *TaskStore) Create(req *http.Request) (*models.Task, error) {
	ctx := req.Context()
	defer ctx.Done()
	return nil, errors.New("Not Implemented")
}

func (store *TaskStore) Update(req *http.Request) (*models.Task, error) {
	ctx := req.Context()
	defer ctx.Done()
	return nil, errors.New("Not Implemented")
}

func (store *TaskStore) Delete(req *http.Request) (*models.Task, error) {
	ctx := req.Context()
	defer ctx.Done()
	return nil, errors.New("Not Implemented")
}

// Get last date by parsing url query parameter "lastDays",
// by default last date is -100 days from now
func getLastDate(query url.Values) int64 {
	lastDays := query.Get("lastDays")
	days, err := strconv.Atoi(lastDays)
	if err != nil {
		days = -100
	} else {
		days = -days
	}
	lastUnix := time.Now().AddDate(0, 0, days).Unix()
	return lastUnix * 1000
}
