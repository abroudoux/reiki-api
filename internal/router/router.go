package router

import (
	"log"
	"time"

	"github.com/abroudoux/reiki-api/internal/database"
	services "github.com/abroudoux/reiki-api/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(config))

	err := database.CreateTableSessions()

	if err != nil {
		log.Fatalf("Failed to create sessions table: %v", err)
	}

	err = database.CreateTableMessages()

	if err != nil {
		log.Fatalf("Failed to create messages table: %v", err)
	}

	router.GET("/hello", services.HelloWorld)

	router.GET("/sessions", services.GetSessions)
	router.POST("/sessions", services.PostSession)
	router.DELETE("/sessions/:id", services.DeleteSession)

	router.GET("/messages", services.GetMessages)
	router.POST("/messages", services.PostMessage)
	router.DELETE("/messages/:id", services.DeleteMessage)

	router.Run(":8080")
}
