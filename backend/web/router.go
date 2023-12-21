package web

import (
	"backend/web/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HttpRouter() {

	router := mux.NewRouter()

	//router.HandleFunc("/auth", controllers.BasicAuthMiddleware).Methods("GET")

	router.HandleFunc("/api/bugs", controllers.GetBugs).Methods("GET")
	router.HandleFunc("/api/features", controllers.GetFeatures).Methods("GET")
	router.HandleFunc("/api/all", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/assigned", controllers.GetAllAssigned).Methods("GET")
	router.HandleFunc("/api/unassigned", controllers.GetAllNotAssigned).Methods("GET")
	router.HandleFunc("/api/createTask", controllers.SaveTask).Methods("POST")
	router.HandleFunc("/api/updateTask", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/deleteTask", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/api/assignTask", controllers.AssignTask).Methods("POST")
	router.HandleFunc("/api/getUsersTasks", controllers.GetUsersTasks).Methods("POST")
	router.HandleFunc("/api/getReportedTasks", controllers.GetReportedTasks).Methods("POST")

	handler := cors.Default().Handler(router)

	http.ListenAndServe(":8080", handler)
}
