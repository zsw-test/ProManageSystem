package main

import (
	"ProManageSystem/api/complaint"
	"ProManageSystem/api/expressage"
	"ProManageSystem/api/manager"
	"ProManageSystem/api/owner"
	"ProManageSystem/api/park"
	"ProManageSystem/api/repair"
	"ProManageSystem/conf"
	"ProManageSystem/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	router := gin.Default()

	//快件
	router.GET("api/expressagetotal", expressage.ExpressageGetTotal)
	router.GET("api/expressagepage", expressage.ComplaintGetPage)

	router.GET("api/expressage/:id", expressage.ExpressageGet)
	router.POST("api/expressage", expressage.ExpressageCreate)
	router.DELETE("api/expressage/:id", expressage.ExpressageDelete)
	router.PUT("api/expressage/:id", expressage.ExpressageSave)

	//维修
	router.GET("api/repairtotal", repair.RepairGetTotal)
	router.GET("api/repairpage", repair.ComplaintGetPage)

	router.GET("api/repair/:id", repair.RepairGet)
	router.POST("api/repair", repair.RepairCreate)
	router.DELETE("api/repair/:id", repair.RepairDelete)
	router.PUT("api/repair/:id", repair.RepairSave)

	//投诉
	router.GET("api/complainttotal", complaint.ComplaintGetTotal)
	router.GET("api/complaintpage", complaint.ComplaintGetPage)

	router.GET("api/complaint/:id", complaint.ComplaintGet)
	router.POST("api/complaint", complaint.ComplaintCreate)
	router.DELETE("api/complaint/:id", complaint.ComplaintDelete)
	router.PUT("api/complaint/:id", complaint.ComplaintSave)

	//停车
	router.POST("api/parkinfo/:carnumber", park.ParkInfoIn)
	router.GET("api/parkinfo/:carnumber", park.ParkInfoOut)
	router.GET("api/parkinfopage", park.ParkInfoGetPage)
	router.GET("api/parkinfototal", park.ParkInfoGetTotal)
	router.DELETE("api/parkinfo/:carnumber", park.ParkInfoDelete)

	router.POST("api/carinfo", park.CarInfoBuy)

	router.POST("api/park", park.ParkCreate)
	router.GET("api/parkpage", park.ParkGetPage)
	router.GET("api/park/:id", park.ParkGet)
	router.DELETE("api/park/:id", park.ParkDelete)
	router.PUT("api/park/:id", park.ParkSave)

	//用户
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
	//管理员
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
