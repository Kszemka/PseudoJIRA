package db

import (
	"backend/config"
	models "backend/entities"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

const GET_TASKS_QUERY = `MATCH (:CATEGORY {name: $category})<-[:BELONGS_TO]-(task:TASK) RETURN task;`
const GET_ASSIGNED_QUERY = `MATCH (task:TASK) WHERE %s (:USER)-[:ASSIGN]->(task) RETURN task;`
const CREATE_QUERY = `MATCH  (taskCategory:CATEGORY {name: $category}) CREATE (t:TASK {name: $name, description: $description})-[:BELONGS_TO]->(taskCategory);`
const MATCH_USER = `MATCH (user:USER {username: $username}), (task:TASK {name: $name}) CREATE (user)-[:%s]->(task);`
const UPDATE_TASK_QUERY = `MATCH (t:TASK) WHERE t.name = $taskName SET t.name = $name, t.description = $description RETURN t;`
const DELETE_TASK_QUERY = `MATCH (t:TASK)-[r]-() WHERE t.name = $name DELETE r, t`
const GET_USERS_QUERY = `MATCH (user:USER) RETURN user;`
const GET_TASKS_FOR_USER = `MATCH (user:USER {username: $username})-[:%s]->(task:TASK) return task`
const UNASSIGN_TASK = `MATCH (user:USER {username: $username})-[assign:ASSIGN]->(task:TASK {name:$name})
						WHERE (user)-[:ASSIGN]->(task) DELETE assign;`

func RunCreateQuery(con *config.Connection, task models.Task, user models.User) (any, error) {
	session := con.Driver.NewSession(con.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	result, err := session.ExecuteWrite(con.Ctx, func(transaction neo4j.ManagedTransaction) (interface{}, error) {
		result, err := transaction.Run(con.Ctx, CREATE_QUERY, map[string]interface{}{"name": task.Name, "description": task.Description, "category": task.Category})
		return result, err
	})
	if err != nil {
		log.Println("ERROR with Create Node")
		return result, err
	}

	result, err = session.ExecuteWrite(con.Ctx, func(transaction neo4j.ManagedTransaction) (interface{}, error) {
		result, err := transaction.Run(con.Ctx, fmt.Sprintf(MATCH_USER, "REPORT"), map[string]interface{}{"name": task.Name, "username": user.Username})
		return result, err
	})
	if err != nil {
		log.Println("ERROR with Create Relation")
	}

	return result, err
}

func RunQuery(con *config.Connection, QUERY string, params map[string]any) (neo4j.ResultWithContext, error) {
	session := con.Driver.NewSession(con.Ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	result, err := session.Run(
		con.Ctx,
		QUERY,
		params,
	)
	if err != nil {
		log.Println("Wrong query on database")
		return nil, err
	}

	return result, err
}
