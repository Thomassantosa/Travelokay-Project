package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/Travelokay-Project/models"
)

func SendSuccessResponse(w http.ResponseWriter) {

	// success response
	response := models.MessageResponse{}
	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, errCode int) {
	// var response MessageResponse
	response := models.MessageResponse{}
	response.Status = errCode

	switch errCode {
	case 500:
		response.Message = "Internal server error"
	case 400:
		response.Message = "Bad Request"
	default:
		response.Message = "Undeclared Error"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(response)
}
