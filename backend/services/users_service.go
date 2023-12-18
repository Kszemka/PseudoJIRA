package services

import (
	"backend/config"
	database "backend/config/db"
	models "backend/entities"
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

func ExecuteGetUsersQuery() ([]models.User, error) {
	db := config.GetDbConnection()
	var users []models.User

	result, err := database.RunQuery(db, database.GET_USERS_QUERY, nil)
	if err != nil {
		return nil, err
	}

	var user models.User

	for result.Next(db.Ctx) {
		userRecord := result.Record()
		u, _ := userRecord.Get("user")
		result, _ := json.Marshal(u.(neo4j.Node).Props)
		json.Unmarshal(result, &user)

		users = append(users, user)
	}

	return users, nil
}

func ExecuteAssignTaskToUser(params models.TaskWithAssignedUser) error {
	db := config.GetDbConnection()
	data := map[string]any{}
	data["username"] = params.User.Username
	data["name"] = params.Task.Name

	log.Println(data)
	_, err := database.RunQuery(db, fmt.Sprintf(database.MATCH_USER, "ASSIGN"), data)

	return err
}

func ExecuteShowUsersTasks(user models.User, option string) ([]models.Task, error) {
	db := config.GetDbConnection()
	var tasks []models.Task

	data := map[string]any{}
	data["username"] = user.Username

	result, err := database.RunQuery(db, fmt.Sprintf(database.GET_TASKS_FOR_USER, option), data)
	if err != nil {
		return nil, err
	}

	var task models.Task

	for result.Next(db.Ctx) {
		userRecord := result.Record()
		t, _ := userRecord.Get("task")
		result, _ := json.Marshal(t.(neo4j.Node).Props)
		json.Unmarshal(result, &task)

		tasks = append(tasks, task)
	}

	return tasks, nil
}
