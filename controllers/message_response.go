package controllers

import (
	"encoding/json"
	"log"
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
	case 204:
		response.Message = "No Content"
	case 500:
		response.Message = "Internal server error"
	case 400:
		response.Message = "Bad Request"
	case 401:
		response.Message = "UnAuthorized Access"
	default:
		response.Message = "Undeclared Error"
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(response)
}

func SendMessageOnlyResponse(w http.ResponseWriter, message string) {

	// success response
	response := models.MessageResponse{}
	response.Status = 200
	response.Message = message

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func PrintError(err error) {
	log.Println(string("\033[31m"), "(ERROR)\t", string("\033[0m"), err)
}

func PrintSuccess(text string) {
	log.Println(string("\033[32m"), "(SUCCESS)\t", string("\033[0m"), text)
}
