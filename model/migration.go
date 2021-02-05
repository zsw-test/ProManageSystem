package model

import (
	"ProManageSystem/DB"
	"ProManageSystem/model/manager"
	"ProManageSystem/model/owner"
	"ProManageSystem/model/parkmodel"
)

//数据迁移
func Migration() {
	DB.Mysqldb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
		AutoMigrate(&owner.Owner{}).
		AutoMigrate(&manager.Manager{}).
		AutoMigrate(&parkmodel.Park{}).
		AutoMigrate(&parkmodel.ParkInfo{}).
		AutoMigrate(&parkmodel.CarInfo{})

}
