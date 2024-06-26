package router

import (
	"time"

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

	router.GET("/hello", services.HelloWorld)
	router.GET("/sessions", services.GetSessions)
	router.POST("/sessions", services.CreateSession)
	router.GET("/messages", services.GetMessages)
	router.POST("/messages", services.CreateMessage)

	router.Run(":8080")
}
