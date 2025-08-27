package dtos

type ProblemResponse struct {
	Id		    string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Editorial   string `json:"editorial"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}