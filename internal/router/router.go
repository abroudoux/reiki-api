package router

import (
	services "github.com/abroudoux/reiki-api/internal/services"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
    router := gin.Default()

	router.GET("/hello", services.HelloWorld) 
	router.GET("/sessions", services.GetSessions)
	router.POST("/sessions", services.CreateSession)

    router.Run(":8080")
}