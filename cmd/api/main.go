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

	router.InitRouter()
}
