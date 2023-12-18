package controllers

import (
	"backend/config"
	models "backend/entities"
	"backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// GetBugs /bugs endpoint
func GetBugs(w http.ResponseWriter, request *http.Request) {
	tasks, err := services.ExecuteGetBugsQuery()
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, tasks)
}

// GetFeatures /features endpoint
func GetFeatures(w http.ResponseWriter, request *http.Request) {
	tasks, err := services.ExecuteGetFeaturesQuery()
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, tasks)
}

// GetAllTasks /all endpoint
func GetAllTasks(w http.ResponseWriter, request *http.Request) {
	tasks1, err := services.ExecuteGetFeaturesQuery()
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	tasks2, err := services.ExecuteGetBugsQuery()
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, append(tasks1, tasks2...))
}

// GetAllAssigned /all endpoint

func GetAllAssigned(w http.ResponseWriter, request *http.Request) {
	tasks, err := services.ExecuteGetAllAssignedQuery()
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, tasks)

}

// GetAllNotAssigned /all endpoint

func GetAllNotAssigned(w http.ResponseWriter, request *http.Request) {
	tasks, err := services.ExecuteGetAllNotAssignedQuery()
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, tasks)

}

func SaveTask(w http.ResponseWriter, request *http.Request) {
	var data models.TaskWithAssignedUser

	err := json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.ExecuteSaveQuery(data)
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, "TASK CREATED")

}

func UpdateTask(w http.ResponseWriter, request *http.Request) {
	var data map[string]any

	err := json.NewDecoder(request.Body).Decode(&data)
	log.Println(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.ExecuteUpdateQuery(data)
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, "TASK UPDATED")

}

func DeleteTask(w http.ResponseWriter, request *http.Request) {
	var task models.Task

	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.ExecuteDeleteQuery(task)
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, "TASK DELETED")

}
