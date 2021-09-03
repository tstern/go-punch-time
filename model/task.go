package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Start int64              `json:"start"`
	End   int64              `json:"end"`
	Text  string             `json:"text"`
	Type  string             `json:"type"`
}
