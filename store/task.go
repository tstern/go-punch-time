package store

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/sternth/go-punch-time/model"
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

func (store *TaskStore) Find(req *http.Request) (*model.Task, error) {
	ctx := req.Context()
	defer ctx.Done()
	return nil, errors.New("Not Implemented")
}

func (store *TaskStore) FindAll(ctx context.Context, lastDays int) ([]model.Task, error) {
	lastDate := time.Now().AddDate(0, 0, -lastDays).Unix() * 1000
	filter := bson.M{"start": bson.M{"$gte": lastDate}}

	cursor, err := store.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	if tasks == nil {
		tasks = []model.Task{}
	}

	return tasks, nil
}

func (store *TaskStore) Create(req *http.Request) (*model.Task, error) {
	ctx := req.Context()
	defer ctx.Done()
	return nil, errors.New("Not Implemented")
}

func (store *TaskStore) Update(req *http.Request) (*model.Task, error) {
	ctx := req.Context()
	defer ctx.Done()
	return nil, errors.New("Not Implemented")
}

func (store *TaskStore) Delete(req *http.Request) (*model.Task, error) {
	ctx := req.Context()
	defer ctx.Done()
	return nil, errors.New("Not Implemented")
}
