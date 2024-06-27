package services

import (
	"net/http"

	"github.com/abroudoux/reiki-api/internal/database"
	"github.com/abroudoux/reiki-api/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetMessage(c *gin.Context) {
	id := c.Param("id")
	message, err := database.GetMessage(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

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

	if message.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "FirstName is required",
		})
		return
	}

	if message.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "LastName is required",
		})
		return
	}

	if message.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is required",
		})
		return
	}

	if message.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Message is required",
		})
		return
	}

	if message.Date == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Date is required",
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
