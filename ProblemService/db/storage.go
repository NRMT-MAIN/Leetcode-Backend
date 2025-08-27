package db

import "leetcode/db/repositories"

type Storage struct {
	ProblemRepository repositories.ProblemRepository
}

func NewStorage() *Storage {
	return &Storage{
		ProblemRepository: &repositories.ProblemRepositoryImpl{},
	}
}
