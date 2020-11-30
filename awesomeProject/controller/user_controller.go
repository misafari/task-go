package controller

import (
	"../model"
	"../service"
	"github.com/gin-gonic/gin"
)

type UserPersistRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
}

func UserSave(c *gin.Context) {
	var requestBody UserPersistRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: requestBody.Username,
		Name:     requestBody.Name,
		Surname:  requestBody.Surname,
	}

	service.UserInsert(&user)

	c.JSON(201, gin.H{"data": true})
}

func UsersFindALl(c *gin.Context) {
	var users []model.User
	service.UsersFindAll(&users)

	c.JSON(200, gin.H{"data": users})
}

func UserFindByUsername(c *gin.Context) {
	var user model.User
	err := service.UserFindByUsername(c.Param("username"), &user)

	if err != nil {
		c.JSON(404, gin.H{"error": "user not found."})
		return
	}

	c.JSON(200, gin.H{"data": user})
}