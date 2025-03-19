package main

import (
	"hotel-database-design/config"
	"hotel-database-design/server"
)

func getMainServer() *server.Server {
	return &server.Server{
		PSQL_CONNECTION: config.PSQL_CONNECTION,
		PORT:            config.PORT,
	}
}

func main() {
	s := getMainServer()
	s.Start()
}
