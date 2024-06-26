package services

import (
	"net/http"

	"github.com/abroudoux/reiki-api/internal/database"
	"github.com/abroudoux/reiki-api/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetSessions(c *gin.Context) {
	sessions, err := database.GetSessions()

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

func PostSession(c *gin.Context) {
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
	err := database.PostSession(session)

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

func DeleteSession(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteSession(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Session deleted",
	})
}