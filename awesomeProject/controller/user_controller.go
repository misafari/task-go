package controller

import (
	"../codeConst"
	"../model"
	"../service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserPersistRequest struct {
	Username string `json:"username" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
}

func UserSave(c *gin.Context) {
	var requestBody UserPersistRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{codeConst.Error: err.Error()})
		return
	}

	user := model.User{
		Username: requestBody.Username,
		Name:     requestBody.Name,
		Surname:  requestBody.Surname,
	}

	service.UserInsert(&user)

	c.JSON(http.StatusCreated, gin.H{codeConst.Data: true})
}

func UsersFindALl(c *gin.Context) {
	var users []model.User
	service.UsersFindAll(&users)

	c.JSON(http.StatusOK, gin.H{codeConst.Data: users})
}

func UserFindByUsername(c *gin.Context) {
	var user model.User
	err := service.UserFindByUsername(c.Param("username"), &user)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{codeConst.Error: codeConst.EntityNotFound("user")})
		return
	}

	c.JSON(http.StatusOK, gin.H{codeConst.Data: user})
}
