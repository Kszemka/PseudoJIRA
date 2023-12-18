package main

import (
	"backend/config"
	"backend/web"
)

func main() {
	config.InitDbConnection()
	web.HttpRouter()
	defer config.CloseDbConn()
}
