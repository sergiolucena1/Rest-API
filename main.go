package main

import (
	"github.com/sergiolucena1/database"
	"github.com/sergiolucena1/routes/server"
)

func main() {
	database.StartDB()

	server := server.NewServer() // criando o server

	server.Run()
}
//go mod init githu.com/sssss
//go mod tidy