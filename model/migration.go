package model

import "ProManageSystem/DB"

//数据迁移
func Migration() {
	DB.Mysqldb.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").
		AutoMigrate(&User{}).
		AutoMigrate(&Owner{}).
		AutoMigrate(&Manager{})

}
