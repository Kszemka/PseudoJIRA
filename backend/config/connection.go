package config

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Connection struct {
	Driver neo4j.DriverWithContext
	Ctx    context.Context
}

type Configuration struct {
	URI      string `json:"URI"`
	Database string `json:"Database"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func NewDbConnection(driver neo4j.DriverWithContext, ctx context.Context) *Connection {
	return &Connection{Driver: driver, Ctx: ctx}
}
