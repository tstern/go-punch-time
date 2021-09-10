package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Start int64              `bson:"start" json:"start"`
	End   int64              `bson:"end" json:"end"`
	Text  string             `bson:"text" json:"text"`
	Type  string             `bson:"type" json:"type"`
}
