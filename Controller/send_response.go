package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	m "uts_pbp/Model"
)

func sendErrorResponse(w http.ResponseWriter, msg string) {
	var response m.AccountResponse
	response.Status = 400
	response.Message = msg
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendAccountSuccessResponse(w http.ResponseWriter, data []m.Account, message string) {
	var response m.AccountsResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Successfully Launched")
	json.NewEncoder(w).Encode(response)
}

func sendRoomsSuccessResponse(w http.ResponseWriter, data []m.Room, message string) {
	var response m.RoomsResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Successfully Launched")
	json.NewEncoder(w).Encode(response)
}

func sendRoomDescSuccessResponse(w http.ResponseWriter, data m.RoomDesc, message string) {
	var response m.RoomDescResponse
	response.Status = 200
	response.Message = message
	response.Data = data
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Successfully Launched")
	json.NewEncoder(w).Encode(response)
}
