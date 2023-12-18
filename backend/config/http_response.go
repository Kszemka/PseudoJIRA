package config

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteHttpResponse(w http.ResponseWriter, output any) {
	result, err := json.Marshal(output)
	if err != nil {
		log.Printf("Error converting tasks to JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
