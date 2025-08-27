package routers

import (
	"leetcode/controllers"

	"github.com/go-chi/chi/v5"
)

type ProblemRouter struct {
	problemController *controllers.ProblemController
}

func NewProblemRouter(_problemController *controllers.ProblemController) Router {
	return &ProblemRouter{
		problemController: _problemController,
	}
}

func (pr *ProblemRouter) Register(r chi.Router) {
	r.Post("/problems", pr.problemController.CreateProblem)
	r.Get("/problems/{id}", pr.problemController.GetProblem)
	r.Put("/problems/{id}", pr.problemController.UpdateProblem)
	r.Delete("/problems/{id}", pr.problemController.DeleteProblem)
	r.Get("/problems", pr.problemController.GetAllProblems)
	r.Get("/problems/search/{query}", pr.problemController.SearchProblems)
}
