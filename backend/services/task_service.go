package services

import (
	"backend/config"
	database "backend/config/db"
	models "backend/entities"
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func ExecuteGetAllAssignedQuery() (any, error) {
	db := config.GetDbConnection()

	result, err := database.RunQuery(db, fmt.Sprintf(database.GET_ASSIGNED_QUERY, ""), nil)
	if err != nil {
		return nil, err
	}
	var task models.Task
	tasks := []models.Task{}

	for result.Next(db.Ctx) {
		mergedRecord := result.Record()
		t, _ := mergedRecord.Get("task")
		result, _ := json.Marshal(t.(neo4j.Node).Props)
		json.Unmarshal(result, &task)

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func ExecuteGetAllNotAssignedQuery() (any, error) {
	db := config.GetDbConnection()

	result, err := database.RunQuery(db, fmt.Sprintf(database.GET_ASSIGNED_QUERY, "NOT"), nil)
	if err != nil {
		return nil, err
	}
	var task models.Task
	tasks := []models.Task{}

	for result.Next(db.Ctx) {
		mergedRecord := result.Record()
		t, _ := mergedRecord.Get("task")
		result, _ := json.Marshal(t.(neo4j.Node).Props)
		json.Unmarshal(result, &task)

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func ExecuteSaveQuery(data models.TaskWithAssignedUser) error {
	db := config.GetDbConnection()

	_, err := database.RunCreateQuery(db, data.Task, data.User)

	return err
}

func ExecuteUpdateQuery(params map[string]any) error {
	db := config.GetDbConnection()

	_, err := database.RunQuery(db, database.UPDATE_TASK_QUERY, params)

	return err
}

func ExecuteDeleteQuery(task models.Task) error {
	db := config.GetDbConnection()
	var data map[string]any
	jTask, _ := json.Marshal(task)
	json.Unmarshal(jTask, &data)

	_, err := database.RunQuery(db, database.DELETE_TASK_QUERY, data)

	return err
}
