package dtos


type SubmissionResponse struct {
	Id          string     `json:"id" bson:"_id"`
	ProblemId   string     `json:"problem_id" bson:"problem_id"`
	Code        string     `json:"code" bson:"code"`
	Status      string     `json:"status" bson:"status"`
	Language    string     `json:"language" bson:"language"`
	CreatedAt   string     `json:"created_at"`
	UpdatedAt   string     `json:"updated_at"`
}

type CreateSubmissionRequest struct {
	ProblemId   string     `json:"problem_id" bson:"problem_id"`
	Code        string     `json:"code" bson:"code"`
	Language    string     `json:"language" bson:"language"`
}