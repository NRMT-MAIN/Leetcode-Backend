package controllers

import (
	"leetcode/models"
	"leetcode/services"
	"leetcode/utils"
	"net/http"
)

type ProblemController struct {
	ProblemService services.ProblemService
}

func NewProblemController(problemService services.ProblemService) *ProblemController {
	return &ProblemController{
		ProblemService: problemService,
	}
}


func (pc *ProblemController) CreateProblem(w http.ResponseWriter, r *http.Request) {
	var problem models.Problem
	if err := utils.ReadJSONRequest(r, &problem); err != nil {
		utils.WriteErrorJSONResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdProblem, err := pc.ProblemService.CreateProblem(&problem)
	if err != nil {
		utils.WriteErrorJSONResponse(w, "Error creating problem", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessJSONResponse(w, "Problem created successfully", http.StatusCreated, createdProblem)
}

func (pc *ProblemController) GetProblem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	problem, err := pc.ProblemService.GetProblem(id)
	if err != nil {
		utils.WriteErrorJSONResponse(w, "Error fetching problem", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessJSONResponse(w, "Problem fetched successfully", http.StatusOK, problem)
}

func (pc *ProblemController) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var problem models.Problem
	if err := utils.ReadJSONRequest(r, &problem); err != nil {
		utils.WriteErrorJSONResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	updatedProblem, err := pc.ProblemService.UpdateProblem(id, &problem)
	if err != nil {
		utils.WriteErrorJSONResponse(w, "Error updating problem", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessJSONResponse(w, "Problem updated successfully", http.StatusOK, updatedProblem)
}

func (pc *ProblemController) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := pc.ProblemService.DeleteProblem(id); err != nil {
		utils.WriteErrorJSONResponse(w, "Error deleting problem", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessJSONResponse(w, "Problem deleted successfully", http.StatusOK, nil)
}

func (pc *ProblemController) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	problems, err := pc.ProblemService.GetAllProblems()
	if err != nil {
		utils.WriteErrorJSONResponse(w, "Error fetching problems", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessJSONResponse(w, "Problems fetched successfully", http.StatusOK, problems)
}

func (pc *ProblemController) SearchProblems(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	problems, err := pc.ProblemService.SearchProblem(query)
	if err != nil {
		utils.WriteErrorJSONResponse(w, "Error searching problems", http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessJSONResponse(w, "Problems fetched successfully", http.StatusOK, problems)
}