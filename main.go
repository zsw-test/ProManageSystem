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
	"ProManageSystem/middleware/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	conf.Init()
	router := gin.Default()
	router.Use(cors.Cors())

	//停车
	//买停车卡(用户)
	router.POST("api/carinfo", park.CarInfoBuy)
	//查看车辆是否有月卡
	router.GET("api/carinfo", park.CarinfoGet)

	//查看所有停在物业的车  和月卡车
	router.GET("api/carinfopage", park.CarinfoGetPage)
	router.GET("api/carinfototal", park.CarinfoGetPage)

	router.POST("api/parkinfo/:carnumber", park.ParkInfoIn)
	router.GET("api/parkinfo/:carnumber", park.ParkInfoOut)
	router.GET("api/parkinfopage", park.ParkInfoGetPage)
	router.GET("api/parkinfototal", park.ParkInfoGetTotal)
	router.DELETE("api/parkinfo/:carnumber", park.ParkInfoDelete)

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
	ownerapi.GET("/:id/total", owner.OwnerGetTotal)
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
	//投诉
	ownerapi.GET("/:id/complaint/:complaintid", complaint.ComplaintGet)
	ownerapi.POST("/:id/complaint", complaint.ComplaintCreate)
	ownerapi.DELETE("/:id/complaint/:complaintid", complaint.ComplaintDelete)
	ownerapi.PUT("/:id/complaint/:complaintid", complaint.ComplaintSave)

	//维修
	ownerapi.GET("/:id/repair/:repairid", repair.RepairGet)
	ownerapi.POST("/:id/repair", repair.RepairCreate)
	ownerapi.DELETE("/:id/repair/:repairid", repair.RepairDelete)
	ownerapi.PUT("/:id/repair/:repairid", repair.RepairSave)

	//快件
	ownerapi.GET("/:id/expressage/:expressageid", expressage.ExpressageGet)
	ownerapi.POST("/:id/expressage", expressage.ExpressageCreate)
	ownerapi.DELETE("/:id/expressage/:expressageid", expressage.ExpressageDelete)
	ownerapi.PUT("/:id/expressage/:expressageid", expressage.ExpressageSave)

	//}

	//管理员
	router.POST("/api/managerlogin", manager.ManagerLogin)
	router.POST("/api/managerregister", manager.ManagerRegister)
	managerapi := router.Group("/api/manager")
	// managerapi.Use(jwt.JWTmanager())
	// {
	managerapi.GET("/:id/page", manager.ManagerGetPage)
	managerapi.GET("/:id/total", manager.ManagerGetTotal)
	managerapi.GET("/:id", manager.ManagerGet)
	managerapi.PUT("/:id", manager.ManagerSave)
	managerapi.DELETE("/:id", manager.ManagerDelete)

	//总体费用查看和记录查询 （管理员）
	managerapi.PUT("/:id/charge/:houseid", charge.ChargeSave)
	managerapi.GET("/:id/chargetotal", charge.ChargeGetTotal)
	managerapi.GET("/:id/chargepage", charge.ChargeGetPage)

	managerapi.GET("/:id/chargerecordpage", charge.ChargeRecordGetPage)
	managerapi.GET("/:id/chargerecordtotal", charge.ChargeRecordGetTotal)

	//查看投诉
	managerapi.GET("/:id/complainttotal", complaint.ComplaintGetTotal)
	managerapi.GET("/:id/complaintpage", complaint.ComplaintGetPage)

	//查看维修
	managerapi.GET("/:id/repairtotal", repair.RepairGetTotal)
	managerapi.GET("/:id/repairpage", repair.ComplaintGetPage)

	//查看All快件
	managerapi.GET("/:id/expressagetotal", expressage.ExpressageGetTotal)
	managerapi.GET("/:id/expressagepage", expressage.ComplaintGetPage)
	//录入快件
	managerapi.POST("/:id/expressage", expressage.ExpressageCreate)

	// }
	//数据准备
	//model.PrepareHouse()
	//model.PrepareUsers()
	router.Run(":31717")
}
