package parkmodel

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

//停车信息（收费，出入时间）
type ParkInfo struct {
	gorm.Model
	//车辆号码
	CarNumber string `gorm:"column:carnumber;type:varchar(20);unique_index:carnumber"`
	//费用  根据当前时间和创建时间来算出总费用
	Fee int
}

func (parkinfo *ParkInfo) Create() error {
	err := DB.Mysqldb.Create(parkinfo).Error
	return err
}

func (parkinfo *ParkInfo) Save() error {
	err := DB.Mysqldb.Save(parkinfo).Error
	return err
}

func (parkinfo *ParkInfo) Delete() error {
	err := DB.Mysqldb.Unscoped().Delete(parkinfo).Error
	return err
}

func GetParkInfobyid(id int) (*ParkInfo, error) {
	var parkinfo = &ParkInfo{}
	err := DB.Mysqldb.Where("id = ?", id).First(&parkinfo).Error
	return parkinfo, err
}
func GetParkInfobyCarnumber(carnumber string) (*ParkInfo, error) {
	var parkinfo = &ParkInfo{}
	err := DB.Mysqldb.Where("carnumber = ?", carnumber).First(&parkinfo).Error
	return parkinfo, err
}

func GetParkInfoPage(pageindex, pagesize int) ([]*ParkInfo, error) {
	parkinfoList := []*ParkInfo{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&parkinfoList).Error
	return parkinfoList, err
}

func GetParkInfoTotal() (int, error) {
	var num int
	err := DB.Mysqldb.Model(&ParkInfo{}).Count(&num).Error
	return num, err
}
