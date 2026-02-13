package main

import (
	"log"
	"notification-engine/internal/config"
)

func main() {
	var err error

	if err = config.Load(); err != nil {
		log.Fatalf("Erro ao carregar as envs: %v", err)
	}

	log.Println("Hello World!")
}
