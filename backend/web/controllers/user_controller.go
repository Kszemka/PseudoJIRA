package controllers

import (
	"backend/config"
	models "backend/entities"
	"backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// GetUsers /users endpoint
func GetAllUsers(w http.ResponseWriter, request *http.Request) {
	users, err := services.ExecuteGetUsersQuery()
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, users)
}

func AssignTask(w http.ResponseWriter, request *http.Request) {
	var data models.TaskWithAssignedUser

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = services.ExecuteAssignTaskToUser(data)
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, "ASSIGNED")

}

func GetUsersTasks(w http.ResponseWriter, request *http.Request) {
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, er := services.ExecuteShowUsersTasks(user, "ASSIGN")
	if er != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, result)
}

func GetReportedTasks(w http.ResponseWriter, request *http.Request) {
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, er := services.ExecuteShowUsersTasks(user, "REPORT")
	if er != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, result)
}
