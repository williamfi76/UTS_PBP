package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	m "uts_pbp/Model"
)

func sendErrorResponse(w http.ResponseWriter, msg string) {
	var response m.UsersResponse
	response.Status = 400
	response.Message = msg
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendUsersSuccessResponse(w http.ResponseWriter, data []m.User, message string) {
	var response m.UsersResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Successfully Launched")
	json.NewEncoder(w).Encode(response)
}
