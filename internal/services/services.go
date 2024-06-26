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
			"error": "Invalid request body",
		})
		return
	}

	if session.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FirstName is required",
		})
		return
	}

	if session.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "LastName is required",
		})
		return
	}

	if session.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is required",
		})
		return
	}

	if session.Date == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Date is required",
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

func GetMessages(c *gin.Context) {
	messages, err := database.ReturnMessages()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": messages,
	})
}

func CreateMessage(c *gin.Context) {
	var message types.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	message.Id = uuid.New().String()
	err := database.AddMessage(message)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Message created",
	})
}
