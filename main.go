package main

import (
	"ProManageSystem/api/charge"
	"ProManageSystem/api/complaint"
	"ProManageSystem/api/expressage"
	"ProManageSystem/api/house"
	"ProManageSystem/api/manager"
	"ProManageSystem/api/owner"
	"ProManageSystem/api/park"
	"ProManageSystem/api/repair"
	"ProManageSystem/api/thirdparty"
	"ProManageSystem/conf"
	"ProManageSystem/middleware/cors"
	"ProManageSystem/middleware/jwt"
	"ProManageSystem/model"
	"flag"

	"github.com/gin-gonic/gin"
)

var p = flag.Bool("p", false, "是否准备数据")

func main() {
	conf.Init()
	flag.Parse()
	//根据命令行参数准备数据
	if *p {
		model.PrepareAll()
	}

	router := gin.Default()
	router.Use(cors.Cors())

	// 百度faceapi
	router.POST("/api/facedetect", thirdparty.FaceDetect)
	router.POST("/api/faceadd", thirdparty.FaceAdd)
	//用户
	router.POST("/api/ownerlogin", owner.OwnerLogin)
	router.POST("/api/ownerregister", owner.OwnerRegister)
	router.GET("/api/qiniutoken", thirdparty.GetQiniuToken)
	router.POST("api/qiniuup", thirdparty.UpPhoto)
	ownerauth := router.Group("/api/ownerauth")
	//测试环境不添加token验证权限都可以进入
	//添加owner的jwt中间件检查
	ownerauth.Use(jwt.JWTowner())
	{
		ownerapi := ownerauth.Group("/owner")
		ownerapi.GET("/:ownerid", owner.OwnerGet)
		ownerapi.PUT("/:ownerid", owner.OwnerSave)
		ownerapi.DELETE("/:ownerid", owner.OwnerDelete)
		//费用缴纳和查询
		ownerauth.GET("/charge/:houseid/water", charge.ChargeGetWater)
		ownerauth.POST("/charge/:houseid/water", charge.ChargeWaterPay)
		ownerauth.GET("/charge/:houseid/electric", charge.ChargeGetElectric)
		ownerauth.POST("/charge/:houseid/electric", charge.ChargeElectricPay)
		ownerauth.GET("/charge/:houseid/gas", charge.ChargeGetGas)
		ownerauth.POST("/charge/:houseid/gas", charge.ChargeGasPay)
		ownerauth.GET("/charge/:houseid/property", charge.ChargeGetProperty)
		ownerauth.POST("/charge/:houseid/property", charge.ChargePropertyPay)
		//总费用查询
		ownerauth.GET("/charge/:houseid", charge.ChargeGet)
		//缴费记录查询
		ownerauth.GET("/chargerecord/:houseid", charge.ChargeRecordGet)

		//我的房屋
		ownerauth.GET("/residenthousepage/:houseid", house.ResidentGetByhouseidPage)
		ownerauth.GET("/residenthousetotal/:houseid", house.ResidentGetByhouseidTotal)
		ownerauth.POST("/resident", house.ResidentAdd)
		ownerauth.DELETE("/resident/:residentid", house.ResidentDel)

		//投诉
		ownerauth.GET("/complaint/:complaintid", complaint.ComplaintGet)
		ownerauth.POST("/complaint", complaint.ComplaintCreate)
		ownerauth.DELETE("/complaint/:complaintid", complaint.ComplaintDelete)
		ownerauth.PUT("/complaint/:complaintid", complaint.ComplaintSave)

		//维修
		ownerauth.GET("/repair/:repairid", repair.RepairGet)
		ownerauth.POST("/repair", repair.RepairCreate)
		ownerauth.DELETE("/repair/:repairid", repair.RepairDelete)
		ownerauth.PUT("/repair/:repairid", repair.RepairSave)
		ownerauth.GET("/repairowner", repair.RepairGetOwner)

		//快件
		ownerauth.GET("/expressage/:expressageid", expressage.ExpressageGet)
		ownerauth.POST("/expressage", expressage.ExpressageCreate)
		ownerauth.DELETE("/expressage/:expressageid", expressage.ExpressageDelete)
		ownerauth.PUT("/expressage/:expressageid", expressage.ExpressageSave)

		//停车
		//买停车卡(用户)
		ownerauth.POST("/carinfo", park.CarInfoBuy)
		//购买车位
		ownerauth.POST("/parkbuy", park.ParkBuy)
		ownerauth.GET("/park", park.ParkGetFreeList)
		//查看车辆是否有月卡
		ownerauth.GET("/carinfo/:carnumber", park.CarinfoGet)
		//用户模拟停车和出库的功能（本来应该是收费站来操作）
		ownerauth.POST("/parkinfo/:carnumber", park.ParkInfoIn)
		ownerauth.GET("/parkinfo/:carnumber", park.ParkInfoOut)
		//支付成功就删除停车信息
		ownerauth.DELETE("/parkinfo/:carnumber", park.ParkInfoDelete)

	}

	//管理员
	router.POST("/api/managerlogin", manager.ManagerLogin)
	router.POST("/api/managerregister", manager.ManagerRegister)
	managerauth := router.Group("/api/managerauth")
	managerauth.Use(jwt.JWTmanager())
	{
		//管理所有的人员  业主和管理员
		managerapi := managerauth.Group("/manager")
		managerapi.GET("/:managerid", manager.ManagerGet)
		managerapi.PUT("/:managerid", manager.ManagerSave)
		managerapi.DELETE("/:managerid", manager.ManagerDelete)

		managerauth.GET("/managerpage", manager.ManagerGetPage)
		managerauth.GET("/managertotal", manager.ManagerGetTotal)
		managerauth.GET("/ownerpage", owner.OwnerGetPage)
		managerauth.GET("/ownertotal", owner.OwnerGetTotal)

		managerauth.DELETE("/owneruntiepark/:ownerid", owner.OwnerUntiepark)
		managerauth.GET("/owner/:ownerid", owner.OwnerGet)
		managerauth.PUT("/owner/:ownerid", owner.OwnerSave)
		managerauth.DELETE("/owner/:ownerid", owner.OwnerDelete)

		//总体费用查看和记录查询 （管理员）
		managerauth.PUT("/charge/:houseid", charge.ChargeSave)
		managerauth.GET("/chargetotal", charge.ChargeGetTotal)
		managerauth.GET("/chargepage", charge.ChargeGetPage)

		managerauth.GET("/chargerecordpage", charge.ChargeRecordGetPage)
		managerauth.GET("/chargerecordtotal", charge.ChargeRecordGetTotal)

		//查看投诉
		managerauth.GET("/complainttotal", complaint.ComplaintGetTotal)
		managerauth.GET("/complaintpage", complaint.ComplaintGetPage)

		//查看维修
		managerauth.GET("/repairtotal", repair.RepairGetTotal)
		managerauth.GET("/repairpage", repair.ComplaintGetPage)
		managerauth.GET("/repairmanager", repair.RepairGetManager)
		//修改维修指派、解决
		managerauth.PUT("/repairdispatch", repair.RepairDispatch)
		managerauth.PUT("/repairresolve", repair.RepairResolve)
		managerauth.DELETE("/repair/:repairid", repair.RepairDelete)

		//查看All快件
		managerauth.GET("/expressagetotal", expressage.ExpressageGetTotal)
		managerauth.GET("/expressagepage", expressage.ComplaintGetPage)
		//录入快件
		managerauth.POST("/expressage", expressage.ExpressageCreate)

		//查看所有停在物业的车  和月卡车（管理员）
		managerauth.GET("/carinfopage", park.CarinfoGetPage)
		managerauth.GET("/carinfototal", park.CarinfoGetTotal)
		managerauth.GET("/parkinfopage", park.ParkInfoGetPage)
		managerauth.GET("/parkinfototal", park.ParkInfoGetTotal)

		managerauth.GET("/parkallcount", park.ParkGetAllCount)

		managerauth.POST("/park", park.ParkCreate)
		managerauth.GET("/parkpage", park.ParkGetPage)
		managerauth.GET("/parktotal", park.ParkGetTotal)
		managerauth.GET("/park/:parkid", park.ParkGet)
		managerauth.DELETE("/park/:parkid", park.ParkDelete)
		managerauth.PUT("/park/:parkid", park.ParkSave)

		managerauth.POST("/house", house.HouseCreate)

		managerauth.GET("/house/:houseid", house.HouseGet)
		managerauth.GET("/housetotal", house.HouseGetTotal)
		managerauth.GET("/housepage", house.HouseGetPage)
		managerauth.DELETE("/house/:houseid", house.HouseDelete)
		managerauth.PUT("/house/:houseid", house.HouseSave)

		managerauth.POST("/resident", house.ResidentAdd)
		managerauth.DELETE("/resident/:residentid", house.ResidentDel)
		managerauth.GET("/resident/:residentid", house.ResidentGet)
		managerauth.PUT("/resident/:residentid", house.ResidentSave)
		managerauth.GET("/residentpage", house.ResidentGetPage)
		managerauth.GET("/residenttotal", house.ResidentGetTotal)
		managerauth.GET("/residenthousepage/:houseid", house.ResidentGetByhouseidPage)
		managerauth.GET("/residenthousetotal/:houseid", house.ResidentGetByhouseidTotal)

	}

	router.Run(":31717")
}
