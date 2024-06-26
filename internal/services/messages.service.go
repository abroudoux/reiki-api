package services

import (
	"net/http"

	"github.com/abroudoux/reiki-api/internal/database"
	"github.com/abroudoux/reiki-api/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMessages(c *gin.Context) {
	messages, err := database.GetMessages()

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

func PostMessage(c *gin.Context) {
	var message types.Message

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	message.Id = uuid.New().String()
	err := database.PostMessage(message)

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

func DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteMessage(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Message deleted",
	})
}
