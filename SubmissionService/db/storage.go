package db

import "Submission_Service/db/repositories"

type Storage struct {
	SubmissionRepository repositories.SubmissionRepository
}

func NewStorage() *Storage {
	return &Storage{
		SubmissionRepository: &repositories.SubmissionRepositoryImpl{},
	}
}