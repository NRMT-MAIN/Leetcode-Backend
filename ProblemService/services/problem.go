package services

import (
	"fmt"
	"leetcode/db/repositories"
	"leetcode/dtos"
	"leetcode/models"
	"leetcode/utils"

	"github.com/microcosm-cc/bluemonday"
)

type ProblemService interface {
	CreateProblem(problem *models.Problem) (*dtos.ProblemResponse, error)
	GetProblem(id string) (*dtos.ProblemResponse, error)
	UpdateProblem(id string, problem *models.Problem) (*dtos.ProblemResponse, error)
	DeleteProblem(id string) error
	SearchProblem(query string) ([]*dtos.ProblemResponse, error)
	GetAllProblems() ([]*dtos.ProblemResponse, error)
}

type ProblemServiceImpl struct {
	ProblemRepository repositories.ProblemRepository
}

func NewProblemService(problemRepository repositories.ProblemRepository) ProblemService {
	return &ProblemServiceImpl{
		ProblemRepository: problemRepository,
	}
}

func (s *ProblemServiceImpl) CreateProblem(problem *models.Problem) (*dtos.ProblemResponse, error) {
	p := bluemonday.NewPolicy()

	// Render the Markdown fields
	problem.Description = utils.RenderMarkdown(problem.Description)
	problem.Editorial = utils.RenderMarkdown(problem.Editorial)

	p.AllowElements("h1", "h2", "h3", "p", "ul", "ol", "li", "strong", "em", "code", "pre")
	p.AllowAttrs("id").OnElements("h1", "h2", "h3", "p", "li")
	sanitizedDesc := string(p.SanitizeBytes([]byte(*problem.Description)))
	problem.Description = &sanitizedDesc
	sanitizedEditorial := string(p.SanitizeBytes([]byte(*problem.Editorial)))
	problem.Editorial = &sanitizedEditorial


	createdProblem, err := s.ProblemRepository.CreateProblem(problem)
	if err != nil {
		fmt.Println("Error creating problem:", err)
		return nil, err
	}
	return createdProblem, nil
}

func (s *ProblemServiceImpl) GetProblem(id string) (*dtos.ProblemResponse, error) {
	problem, err := s.ProblemRepository.GetProblemById(id)
	if err != nil {
		fmt.Println("Error getting problem:", err)
		return nil, err
	}
	return problem , nil
}

func (s *ProblemServiceImpl) UpdateProblem(id string, problem *models.Problem) (*dtos.ProblemResponse, error) {
	p := bluemonday.NewPolicy()
	//Render the Markdown fields
	problem.Description = utils.RenderMarkdown(problem.Description)
	problem.Editorial = utils.RenderMarkdown(problem.Editorial)

	p.AllowElements("h1", "h2", "h3", "p", "ul", "ol", "li", "strong", "em", "code", "pre")
	p.AllowAttrs("id").OnElements("h1", "h2", "h3", "p", "li")
	sanitizedDesc := string(p.SanitizeBytes([]byte(*problem.Description)))
	problem.Description = &sanitizedDesc
	sanitizedEditorial := string(p.SanitizeBytes([]byte(*problem.Editorial)))
	problem.Editorial = &sanitizedEditorial

	updatedProblem, err := s.ProblemRepository.UpdateProblem(id, problem)
	if err != nil {
		fmt.Println("Error updating problem:", err)
		return nil, err
	}
	return updatedProblem, nil
}

func (s *ProblemServiceImpl) DeleteProblem(id string) error {
	err := s.ProblemRepository.DeleteProblem(id)
	if err != nil {
		fmt.Println("Error deleting problem:", err)
		return err
	}
	return nil
}

func (s *ProblemServiceImpl) SearchProblem(query string) ([]*dtos.ProblemResponse, error) {
	problems, err := s.ProblemRepository.SearchProblem(query)
	if err != nil {
		fmt.Println("Error searching problems:", err)
		return nil, err
	}
	return problems, nil
}

func (s *ProblemServiceImpl) GetAllProblems() ([]*dtos.ProblemResponse, error) {
	problems, err := s.ProblemRepository.GetAllProblem()
	if err != nil {
		fmt.Println("Error getting all problems:", err)
		return nil, err
	}
	return problems, nil
}
