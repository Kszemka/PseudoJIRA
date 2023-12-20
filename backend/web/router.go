package web

import (
	"backend/web/controllers"
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func HttpRouter() {

	certFile := os.Getenv("TLS_CERT_PATH")
	keyFile := os.Getenv("TLS_KEY_PATH")

	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	router := mux.NewRouter()
	router.HandleFunc("/bugs", controllers.GetBugs).Methods("GET")
	router.HandleFunc("/features", controllers.GetFeatures).Methods("GET")
	router.HandleFunc("/all", controllers.GetAllTasks).Methods("GET")
	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/assigned", controllers.GetAllAssigned).Methods("GET")
	router.HandleFunc("/unassigned", controllers.GetAllNotAssigned).Methods("GET")
	router.HandleFunc("/createTask", controllers.SaveTask).Methods("POST")
	router.HandleFunc("/updateTask", controllers.UpdateTask).Methods("PUT")
	router.HandleFunc("/deleteTask", controllers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/assignTask", controllers.AssignTask).Methods("POST")
	//router.HandleFunc("/auth", controllers.BasicAuthMiddleware).Methods("GET")
	router.HandleFunc("/getUsersTasks", controllers.GetUsersTasks).Methods("POST")
	router.HandleFunc("/getReportedTasks", controllers.GetReportedTasks).Methods("POST")

	http.Handle("/", corsHandler(router))
	fmt.Println("Server is running on :8443")
	http.ListenAndServeTLS(":8443", certFile, keyFile, nil)}
