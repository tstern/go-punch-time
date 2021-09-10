package store

import (
	"context"
	"fmt"
	"time"

	"github.com/sternth/go-punch-time/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (store *TaskStore) Find(ctx context.Context, id primitive.ObjectID) (*model.Task, error) {
	var task model.Task
	filter := bson.M{"_id": bson.M{"$eq": id}}

	err := store.collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
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

func (store *TaskStore) Create(ctx context.Context, task *model.Task) error {
	task.ID = primitive.NewObjectID()
	_, err := store.collection.InsertOne(ctx, task)
	if err != nil {
		return fmt.Errorf("couldn't add task: %v", err)
	}

	return nil
}

func (store *TaskStore) Update(ctx context.Context, id primitive.ObjectID, task *model.Task) error {
	task.ID = id
	filter := bson.M{"_id": bson.M{"$eq": id}}
	result, err := store.collection.ReplaceOne(ctx, filter, task)
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	if err != nil {
		return fmt.Errorf("couldn't update task with id %s", id)
	}

	return nil
}

func (store *TaskStore) Delete(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{"_id": bson.M{"$eq": id}}
	_, err := store.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("couldn't delete task with id %s", id)
	}

	return nil
}
