package main

import (
	"ProManageSystem/api"
	"ProManageSystem/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/register", api.RegisterHtml)
	router.GET("/login", api.LoginHtml)
	router.POST("/login", api.Login)

	router.POST("/Owner", api.CreateOwner)
	router.GET("/Owner/:username", api.GetOwner)
	router.DELETE("/Owner/:username", api.DeleteOwner)
	router.PUT("/Owner/:username", api.SaveOwner)

	router.POST("/Manager", api.CreateManager)
	router.GET("/Manager/:username", api.GetManager)
	router.DELETE("/Manager/:username", api.DeleteManager)
	router.PUT("/Manager/:username", api.SaveManager)

	router.Run(":31717")
}
