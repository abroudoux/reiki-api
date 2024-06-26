package router

import (
	services "github.com/abroudoux/reiki-api/internal/services"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
    router := gin.Default()

	router.GET("/hello", services.HelloWorld) 

    router.Run(":8080")
}