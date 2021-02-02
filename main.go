package main

import (
	"ProManageSystem/api/manager"
	"ProManageSystem/api/owner"
	"ProManageSystem/conf"
	"ProManageSystem/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	router := gin.Default()
	router.POST("/api/ownerlogin", owner.OwnerLogin)
	router.POST("/api/ownerregister", owner.OwnerRegister)
	ownerauth := router.Group("/api/owner")
	//添加owner的jwt中间件检查
	ownerauth.Use(jwt.JWTowner())
	{
		ownerauth.GET("/:id/page", owner.OwnerGetPage)
		ownerauth.GET("/:id", owner.OwnerGet)
		ownerauth.PUT("/:id", owner.OwnerSave)
		ownerauth.DELETE("/:id", owner.OwnerDelete)
	}

	router.POST("/api/managerlogin", manager.ManagerLogin)
	router.POST("/api/managerregister", manager.ManagerRegister)
	managerauth := router.Group("/api/manager")
	managerauth.Use(jwt.JWTmanager())
	{
		managerauth.GET("/:id/page", manager.ManagerGetPage)
		managerauth.GET("/:id", manager.ManagerGet)
		managerauth.PUT("/:id", manager.ManagerSave)
		managerauth.DELETE("/:id", manager.ManagerDelete)
	}

	router.Run(":31717")
}
