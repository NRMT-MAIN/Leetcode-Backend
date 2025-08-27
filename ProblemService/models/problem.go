package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	Easy   Status = "Easy"
	Medium Status = "Medium"
	Hard   Status = "Hard"
)

type Problem struct {
	Title       string				 `json:"title" bson:"title"`
	Description string                `json:"description" bson:"description"`
	Difficulty  Status                `json:"difficulty" bson:"difficulty"`
	Editorial   string                `json:"editorial" bson:"editorial"`
	CreatedAt   primitive.Timestamp   `json:"created_at" bson:"created_at"`
	UpdatedAt   primitive.Timestamp   `json:"updated_at" bson:"updated_at"`
}