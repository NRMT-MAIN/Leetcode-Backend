package models

import (
	"time"

)

type Status string

const (
	Easy   Status = "Easy"
	Medium Status = "Medium"
	Hard   Status = "Hard"
)

type TestCase struct {
    Input    string `json:"input" bson:"input"`
    Expected string `json:"expected" bson:"expected"`
}


type Problem struct {
	Title       *string				   `json:"title,omitempty" bson:"title"`
	Description *string                `json:"description,omitempty" bson:"description"`
	Difficulty  *Status                `json:"difficulty,omitempty" bson:"difficulty"`
	Editorial   *string                `json:"editorial,omitempty" bson:"editorial"`
	TestCases  *[]TestCase             `json:"test_cases,omitempty" bson:"test_cases"`
	CreatedAt   *time.Time             `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt   *time.Time             `json:"updated_at,omitempty" bson:"updated_at"`
}
