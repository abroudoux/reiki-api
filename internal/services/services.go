package services

import (
	"net/http"

	"github.com/abroudoux/reiki-api/internal/database"
	"github.com/abroudoux/reiki-api/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func GetSessions(c *gin.Context) {
	sessions, err := database.ReturnSessions()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sessions": sessions,
	})
}

func CreateSession(c *gin.Context) {
	var session types.Session

	if err := c.ShouldBindJSON(&session); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	session.Id = uuid.New().String()
	err := database.AddSession(session)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Session created",
	})
}
