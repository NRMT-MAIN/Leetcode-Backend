package dtos

type ProblemResponse struct {
	Id		    string `json:"id" bson:"_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Editorial   string `json:"editorial" bson:"editorial"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}