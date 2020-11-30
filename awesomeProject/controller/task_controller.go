package controller

import (
	"../model"
	"../service"
	"github.com/gin-gonic/gin"
)

type TaskPersistRequest struct {
	Subject       string `json:"subject" binding:"required"`
	Description   string `json:"description"`
	OwnerUsername string `json:"owner"`
}

func TaskSave(c *gin.Context) {
	var requestBody TaskPersistRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	var task = model.Task{
		Subject:     requestBody.Subject,
		Description: requestBody.Description,
	}

	service.TaskPersist(requestBody.OwnerUsername, &task)

	c.JSON(201, gin.H{"data": true})
}
