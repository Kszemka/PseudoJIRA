package config

import (
	"context"
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var db *Connection

var config = Configuration{
	URI:      "neo4j+s://03121ca1.databases.neo4j.io",
	Database: "neo4j",
	Username: "neo4j",
	Password: "FZ8yUeCYq4Kpmm4BTWte2d5KDe3_qw0wdOJuYUcFMXw",
}

func InitDbConnection() *Connection {
	ctx := context.Background()
	var err error

	driver, err := neo4j.NewDriverWithContext(
		config.URI,
		neo4j.BasicAuth(config.Username, config.Password, ""))

	if err != nil {
		log.Fatal(err)
	}
	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		log.Println("FATAL not connected to db")
	}
	fmt.Println("Connected to the Neo4j database")
	db = NewDbConnection(driver, ctx)
	return db

}

func GetDbConnection() *Connection {
	return db
}

func CloseDbConn() {
	db.Driver.Close(db.Ctx)
}
