package main

import (
	"./controller"
	"./service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	service.ConnectToDatabase()

	r.POST("/user/add", controller.UserSave)
	r.GET("/user/byUsername/:username", controller.UserFindByUsername)
	r.GET("/user", controller.UsersFindALl)

	r.POST("/task", controller.TaskSave)

	r.Run()
}