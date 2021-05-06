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
	//时长
	Parktime string
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

func GetParkInfoPage(pageindex, pagesize int, keyword string) ([]ParkInfo, error) {
	ParkInfoList := []ParkInfo{}
	err := DB.Mysqldb.Where("carnumber LIKE ?", "%"+keyword+"%").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&ParkInfoList).Error
	return ParkInfoList, err
}

func GetParkInfoTotal(keyword string) (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&ParkInfo{}).Where("carnumber LIKE ?", "%"+keyword+"%").Count(&count).Error
	return count, err
}
