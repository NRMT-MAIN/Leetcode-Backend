package models

import "time"

type Submission struct {
	ID           *string     `bson:"_id,omitempty"`
	ProblemID    *string     `bson:"problem_id,omitempty"`
	Code         *string     `bson:"code,omitempty"`
	Status       *string     `bson:"status,omitempty"`
	Language     *string     `bson:"language,omitempty"`
	CreatedAt    time.Time   `bson:"created_at,omitempty"`
	UpdatedAt    time.Time   `bson:"updated_at,omitempty"`
}

