package model

import (
	"ProManageSystem/DB"
	"ProManageSystem/model/complaint"
	"ProManageSystem/model/expressage"
	"ProManageSystem/model/manager"
	"ProManageSystem/model/owner"
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/model/repair"
)

//数据迁移
func Migration() {
	DB.Mysqldb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
		AutoMigrate(&owner.Owner{}).
		AutoMigrate(&manager.Manager{}).
		AutoMigrate(&parkmodel.Park{}).
		AutoMigrate(&parkmodel.ParkInfo{}).
		AutoMigrate(&parkmodel.CarInfo{}).
		AutoMigrate(&complaint.Complaint{}).
		AutoMigrate(&expressage.Expressage{}).
		AutoMigrate(&repair.Repair{})

}
