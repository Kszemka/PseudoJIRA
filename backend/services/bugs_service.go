package services

import (
	"backend/config"
	database "backend/config/db"
	models "backend/entities"
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ExecuteGetBugsQuery() ([]models.Task, error) {
	db := config.GetDbConnection()
	tasks := []models.Task{}

	result, err := database.RunQuery(db, database.GET_TASKS_QUERY, map[string]any{"category": "Bug"})
	if err != nil {
		return nil, err
	}

	var record models.Task
	for result.Next(db.Ctx) {
		taskRecord := result.Record()
		t, _ := taskRecord.Get("task")
		result, _ := json.Marshal(t.(neo4j.Node).Props)
		json.Unmarshal(result, &record)
		tasks = append(tasks, record)
	}

	return tasks, nil
}
