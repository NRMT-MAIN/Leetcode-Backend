package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Status string

const (
	Easy   Status = "Easy"
	Medium Status = "Medium"
	Hard   Status = "Hard"
)

type Problem struct {
	Title       *string				 `json:"title,omitempty" bson:"title"`
	Description *string                `json:"description,omitempty" bson:"description"`
	Difficulty  *Status                `json:"difficulty,omitempty" bson:"difficulty"`
	Editorial   *string                `json:"editorial,omitempty" bson:"editorial"`
	CreatedAt   *primitive.Timestamp   `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt   *primitive.Timestamp   `json:"updated_at,omitempty" bson:"updated_at"`
}