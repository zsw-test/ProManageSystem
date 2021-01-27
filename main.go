package main

import (
	"ProManageSystem/api"
	"ProManageSystem/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world!")
	})

	router.POST("/user", api.CreateUser)
	router.Run(":31717")
}
