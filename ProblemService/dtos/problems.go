package dtos

import "leetcode/models"

type ProblemResponse struct {
	Id          string                `json:"id" bson:"_id"`
	Title       string                `json:"title" bson:"title"`
	Description string                `json:"description" bson:"description"`
	Editorial   string                `json:"editorial" bson:"editorial"`
	TestCases   []models.TestCase     `json:"test_cases" bson:"test_cases"`
	CreatedAt   string                `json:"created_at"`
	UpdatedAt   string                `json:"updated_at"`
}