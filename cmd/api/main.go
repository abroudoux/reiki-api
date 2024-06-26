package main

import (
	"log"

	"github.com/abroudoux/reiki-api/internal/database"
	"github.com/abroudoux/reiki-api/internal/router"
)

func main() {
	err := database.CreateTableSessions()

	if err != nil {
		log.Fatalf("Failed to create sessions table: %v", err)
	}

	err = database.CreateTableMessages()

	if err != nil {
		log.Fatalf("Failed to create messages table: %v", err)
	}

	router.InitRouter()
}
