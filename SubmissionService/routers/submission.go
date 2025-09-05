package routers

import (
	"Submission_Service/controllers"
	"github.com/go-chi/chi/v5"
)

type SubmissionRouter struct {
	Controller *controllers.SubmissionController
}

func NewSubmissionRouter(_controller *controllers.SubmissionController) Router {
	return &SubmissionRouter{
		Controller: _controller,
	}
}

func (sr *SubmissionRouter) Register(r chi.Router) {
	r.Post("/submissions", sr.Controller.CreateSubmission)
	r.Get("/submissions/{id}", sr.Controller.GetSubmission)
	r.Put("/submissions/{id}", sr.Controller.UpdateSubmission)
	r.Delete("/submissions/{id}", sr.Controller.DeleteSubmission)
}

