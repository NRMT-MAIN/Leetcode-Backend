package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init(){
	Validator = NewValidator()
}

func NewValidator() *validator.Validate{
	return validator.New(validator.WithRequiredStructEnabled())
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type" , "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func WriteSuccessJSONResponse(w http.ResponseWriter, message string , status int , data any) {
	response := map[string]any{
		"status":  status,
		"message": message,
		"data":    data,
	}
	WriteJSONResponse(w, http.StatusOK, response)
}

func WriteErrorJSONResponse(w http.ResponseWriter, message string , status int) {
	response := map[string]any{
		"status":  status,
		"message": message,
	}
	WriteJSONResponse(w, status, response)
}

func ReadJSONRequest(r *http.Request, result any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(result)
}