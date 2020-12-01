package main

import (
	"./codeConst"
	"./controller"
	"./service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const (
	DefaultAddressPort = ":8081"
	SpecialRoninCode   = 48
)

func main() {
	engine := gin.Default()
	engine.Use(tokenValidation)
	service.ConnectToDatabase()

	engine.POST("/user/add", controller.UserSave)
	engine.GET("/user/byUsername/:username", controller.UserFindByUsername)
	engine.GET("/user", controller.UsersFindALl)

	engine.POST("/task", controller.TaskSave)

	if err := engine.Run(DefaultAddressPort); err != nil {
		fmt.Println("Server Stopped Working since : ", err.Error())
		os.Exit(SpecialRoninCode)
	}
}

func tokenValidation(context *gin.Context) {
	apiToken := context.Request.Header.Get(service.Authorization)
	cleanedToken := apiToken[len(service.Bearer):]

	if apiToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{codeConst.Error: codeConst.TokenIsRequired})
		return
	}

	validationResult, err := service.JWTAuthService().ValidateToken(cleanedToken)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{codeConst.Error: err.Error()})
		return
	}

	if validationResult != nil && validationResult.Valid {
		context.Next()
	} else {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{codeConst.Error: codeConst.TokenNotValid})
	}
}
