package controllers

import (
	"Submission_Service/dtos"
	"Submission_Service/service"
	"Submission_Service/utils"
	"net/http"
)

type SubmissionController struct {
	SubmissionService service.SubmissionService
}

func NewSubmissionController(submissionService service.SubmissionService) *SubmissionController {
	return &SubmissionController{
		SubmissionService: submissionService,
	}
}

func (pc *SubmissionController) CreateSubmission(w http.ResponseWriter, r *http.Request) {
    submission := dtos.NewCreateSubmissionRequest()

    if r.Body == nil {
        utils.WriteJSONResponse(w, http.StatusBadRequest, "Request body is empty")
        return 
    }

    if err := utils.ReadJSONRequest(r, submission); err != nil {
        utils.WriteJSONResponse(w, http.StatusBadRequest, err.Error())
        return
    }

    createdSubmission, err := pc.SubmissionService.CreateSubmission(submission)
    if err != nil {
        utils.WriteJSONResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.WriteJSONResponse(w, http.StatusCreated, createdSubmission)
}

func (pc *SubmissionController) GetSubmission(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	submission, err := pc.SubmissionService.GetSubmissionByID(id)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, submission)
}

func (pc *SubmissionController) UpdateSubmission(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var submission dtos.CreateSubmissionRequest
	if err := utils.ReadJSONRequest(r, &submission); err != nil {
		utils.WriteJSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	updatedSubmission, err := pc.SubmissionService.UpdateSubmission(id, &submission)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, updatedSubmission)
}

func (pc *SubmissionController) DeleteSubmission(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := pc.SubmissionService.DeleteSubmission(id)
	if err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.WriteJSONResponse(w, http.StatusOK, map[string]string{"message": "Submission deleted successfully"})
}
