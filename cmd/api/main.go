package main

import (
	"log"
	"notification-engine/internal/config"
)

func main() {
	var err error

	if err = config.Load(); err != nil {
		log.Fatal("Erro ao carregar as envs")
	}

	log.Println("Hello World!")
}
