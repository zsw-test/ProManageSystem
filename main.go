package main

import (
	"ProManageSystem/api/charge"
	"ProManageSystem/api/complaint"
	"ProManageSystem/api/expressage"
	"ProManageSystem/api/manager"
	"ProManageSystem/api/owner"
	"ProManageSystem/api/park"
	"ProManageSystem/api/repair"
	"ProManageSystem/conf"
	"ProManageSystem/model"

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
	ownerapi := router.Group("/api/owner")
	//测试环境不添加token验证权限都可以进入
	//添加owner的jwt中间件检查
	// ownerapi.Use(jwt.JWTowner())
	// {
	ownerapi.GET("/:id/page", owner.OwnerGetPage)
	ownerapi.GET("/:id", owner.OwnerGet)
	ownerapi.PUT("/:id", owner.OwnerSave)
	ownerapi.DELETE("/:id", owner.OwnerDelete)
	//费用缴纳和查询
	ownerapi.GET("/:id/charge/:houseid/water", charge.ChargeGetWater)
	ownerapi.POST("/:id/charge/:houseid/water", charge.ChargeWaterPay)
	ownerapi.GET("/:id/charge/:houseid/electric", charge.ChargeGetElectric)
	ownerapi.POST("/:id/charge/:houseid/electric", charge.ChargeElectricPay)
	ownerapi.GET("/:id/charge/:houseid/gas", charge.ChargeGetGas)
	ownerapi.POST("/:id/charge/:houseid/gas", charge.ChargeGasPay)
	ownerapi.GET("/:id/charge/:houseid/property", charge.ChargeGetProperty)
	ownerapi.POST("/:id/charge/:houseid/property", charge.ChargePropertyPay)
	//总费用查询
	ownerapi.GET("/:id/charge/:houseid", charge.ChargeGet)
	//缴费记录查询
	ownerapi.GET("/:id/chargerecord/:houseid", charge.ChargeRecordGet)

	//}

	//管理员
	router.POST("/api/managerlogin", manager.ManagerLogin)
	router.POST("/api/managerregister", manager.ManagerRegister)
	managerapi := router.Group("/api/manager")
	// managerapi.Use(jwt.JWTmanager())
	// {
	managerapi.GET("/:id/page", manager.ManagerGetPage)
	managerapi.GET("/:id", manager.ManagerGet)
	managerapi.PUT("/:id", manager.ManagerSave)
	managerapi.DELETE("/:id", manager.ManagerDelete)

	//总体费用查看和记录查询 （管理员）
	managerapi.PUT("/:id/charge/:houseid", charge.ChargeSave)
	managerapi.GET("/:id/chargetotal", charge.ChargeGetTotal)
	managerapi.GET("/:id/chargepage", charge.ChargeGetPage)

	managerapi.GET("/:id/chargerecordpage", charge.ChargeRecordGetPage)
	managerapi.GET("/:id/chargerecordtotal", charge.ChargeRecordGetTotal)

	// }
	//数据准备
	//model.PrepareHouse()
	model.PrepareUsers()
	router.Run(":31717")
}
