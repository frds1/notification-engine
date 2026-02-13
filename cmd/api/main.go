package main

import (
	"log"
	"notification-engine/internal/config"
	"notification-engine/internal/config/database"
	"notification-engine/internal/config/dependencies"
)

func main() {
	var err error

	if err = config.Load(); err != nil {
		log.Fatal("Erro ao carregar as envs")
	}

	dbConn, err := database.OpenConnection(config.GetConfig())
	if err != nil {
		log.Fatal(err)
	}

	app := dependencies.NewContainer(dbConn)

	log.Println("🚀 Server rodando na porta 8080")
	if err := app.Server.Router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
