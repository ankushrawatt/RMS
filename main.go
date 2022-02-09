package main

import (
	"fmt"
	"rmsProject/database"
	"rmsProject/routes"
	"rmsProject/utils"
)

const (
	host     = "localhost"
	port     = "5432"
	dbname   = "rmsProject"
	password = "1234"
	user     = "postgres"
)

func main() {
	err := database.Connection(host, port, dbname, user, password, database.SSLModeDisable)
	utils.CheckError(err)
	fmt.Println("Connection established successfully.....")
	srv := routes.Route()
	connErr := srv.Run(":6666")
	utils.CheckError(connErr)
}
