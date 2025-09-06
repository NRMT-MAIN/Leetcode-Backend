package service

import (
	"Submission_Service/api"
	"Submission_Service/db/repositories"
	"Submission_Service/dtos"
	"fmt"
)

type SubmissionService interface {
	CreateSubmission(submission *dtos.CreateSubmissionRequest) (dtos.SubmissionResponse, error)
	GetSubmissionByID(id string) (*dtos.SubmissionResponse, error)
	UpdateSubmission(id string, submission *dtos.CreateSubmissionRequest) (*dtos.SubmissionResponse, error)
	DeleteSubmission(id string) error
}

type SubmissionServiceImpl struct {
	SubmissionRepository repositories.SubmissionRepository
}

func NewSubmissionService(submissionRepository repositories.SubmissionRepository) SubmissionService {
	return &SubmissionServiceImpl{
		SubmissionRepository: submissionRepository , 
	}
}

func (s *SubmissionServiceImpl) CreateSubmission(submission *dtos.CreateSubmissionRequest) (dtos.SubmissionResponse, error) {
	if submission.Code == nil || submission.Language == nil || submission.ProblemId == nil {
		return dtos.SubmissionResponse{}, fmt.Errorf("invalid submission data")
	}

	resp , err := api.GetProblemById(*submission.ProblemId) ; 
	if err != nil {
		fmt.Println("Error fetching problem details:", err)
		return dtos.SubmissionResponse{} , err
	}
	if string(resp.Id) == "" {
		return dtos.SubmissionResponse{} , fmt.Errorf("problem not found with id: %s" , *submission.ProblemId) 
	}

	fmt.Println("Fetched problem details:", resp.Id , resp.Title) ;

	createdSubmission, err := s.SubmissionRepository.CreateSubmission(submission)
	if err != nil {
		fmt.Println("Error creating submission:", err)
		return dtos.SubmissionResponse{} , err
	}
	return createdSubmission, nil
}

func (s *SubmissionServiceImpl) GetSubmissionByID(id string) (*dtos.SubmissionResponse, error) {
	submission, err := s.SubmissionRepository.GetSubmissionByID(id)
	if err != nil {
		fmt.Println("Error getting submission by ID:", err)
		return nil, err
	}
	return submission, nil			
}

func (s *SubmissionServiceImpl) UpdateSubmission(id string, submission *dtos.CreateSubmissionRequest) (*dtos.SubmissionResponse, error) {
	updatedSubmission, err := s.SubmissionRepository.UpdateSubmission(id, submission)
	if err != nil {
		fmt.Println("Error updating submission:", err)
		return nil, err
	}
	return updatedSubmission, nil
}

func (s *SubmissionServiceImpl) DeleteSubmission(id string) error {
	err := s.SubmissionRepository.DeleteSubmission(id)
	if err != nil {
		fmt.Println("Error deleting submission:", err)
		return err
	}
	return nil
}
