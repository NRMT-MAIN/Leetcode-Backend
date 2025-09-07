package dtos

type Status string

const (
	Pending  Status = "Pending"
	Accepted Status = "Accepted"
	Rejected Status = "Rejected"
)

type SubmissionResponse struct {
	Id          string     `json:"id" bson:"_id"`
	ProblemId   string     `json:"problem_id" bson:"problem_id"`
	Code        string     `json:"code" bson:"code"`
	Status      Status     `json:"status" bson:"status"`
	Language    string     `json:"language" bson:"language"`
	CreatedAt   string     `json:"created_at" bson:"created_at"`
	UpdatedAt   string     `json:"updated_at" bson:"updated_at"`
}

type CreateSubmissionRequest struct {
	ProblemId   *string     `json:"problemId" bson:"problem_id"`
	Code        *string     `json:"code" bson:"code"`
	Language    *string     `json:"language" bson:"language"`
}

func NewCreateSubmissionRequest() *CreateSubmissionRequest {
    return &CreateSubmissionRequest{
        ProblemId:   new(string),
        Code:        new(string),
        Language:    new(string),
    }
}


type SubmmissionJob struct {
	SubmissionId string `json:"submission_id"`
	ProblemId    string `json:"problem_id"`
	Code         string `json:"code"`
	Language     string `json:"language"`
}