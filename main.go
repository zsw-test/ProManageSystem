package main

import (
	"ProManageSystem/api/manager"
	"ProManageSystem/api/owner"
	"ProManageSystem/conf"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	router := gin.Default()

	router.POST("api/owner/login", owner.OwnerLogin)
	router.POST("api/owner/register", owner.OwnerRegister)

	router.POST("api/manager/login", manager.ManagerLogin)
	router.POST("api/manager/register", manager.ManagerRegister)
	router.GET("api/manager/:id", manager.ManagerGet)
	router.PUT("api/manager/:id", manager.ManagerSave)
	router.DELETE("api/manager/:id", manager.ManagerDelete)
	router.Run(":31717")
}
