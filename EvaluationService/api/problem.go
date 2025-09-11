package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TestCase struct {
	Input    string `json:"input"`
	Expected string `json:"expected"`
}

type ProblemResponse struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Editorial   string   `json:"editorial"`
	TestCases   []TestCase `json:"test_cases"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

type problemAPIResponse struct {
	Data    ProblemResponse `json:"data"`
	Message string          `json:"message"`
	Status  int             `json:"status"`
}

func GetProblemById(problemId string) (*ProblemResponse, error) {
	resp, err := http.Get("http://localhost:8080/problems/" + problemId)
	if err != nil {
		fmt.Println("Error fetching problem:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Decode into wrapper
	var apiResp problemAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}

	fmt.Println("Parsed problem:", apiResp.Data)

	return &apiResp.Data, nil
}
