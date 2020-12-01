package controller

import (
	"../codeConst"
	"../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TaskPersistRequest struct {
	Subject       string `json:"subject" binding:"required"`
	Description   string `json:"description"`
	OwnerUsername string `json:"owner"`
}

func TaskSave(c *gin.Context) {
	var requestBody TaskPersistRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{codeConst.Error: err.Error()})
		return
	}

	var task = model.Task{
		Subject:     requestBody.Subject,
		Description: requestBody.Description,
	}

	if err := service.TaskPersist(requestBody.OwnerUsername, &task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{codeConst.Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{codeConst.Data: true})
}
