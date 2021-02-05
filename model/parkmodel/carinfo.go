package parkmodel

import (
	"ProManageSystem/DB"
	"time"

	"github.com/jinzhu/gorm"
)

const (
	TempoaryCar string = "临时车"
	MounthCar   string = "月卡车"
	YearCar     string = "年卡车"
)

type CarInfo struct {
	gorm.Model
	CarNumber string `gorm:"column:carnumber;type:varchar(20);unique_index:carnumber"`
	//车辆状态  是否为临时车、月卡车、年卡车等
	CarType string `gorm:"column:cartype"`
	//到期时间
	ExpireTime time.Time `gorm:"column:expiretime"`
}

func (carinfo *CarInfo) Create() error {
	err := DB.Mysqldb.Create(carinfo).Error
	return err
}

func (carinfo *CarInfo) Save() error {
	err := DB.Mysqldb.Save(carinfo).Error
	return err
}

func (carinfo *CarInfo) Delete() error {
	err := DB.Mysqldb.Unscoped().Delete(carinfo).Error
	return err
}

func GetCarInfobyid(id int) (*CarInfo, error) {
	var carinfo = &CarInfo{}
	err := DB.Mysqldb.Where("id = ?", id).Find(&carinfo).Error
	return carinfo, err
}

func GetCarInfoPage(pageindex, pagesize int) ([]CarInfo, error) {
	carinfolist := []CarInfo{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&carinfolist).Error
	return carinfolist, err
}
func GetCarInfobyCarnumber(carnumber string) (*CarInfo, error) {
	var carinfo = &CarInfo{}
	err := DB.Mysqldb.Where("carnumber = ?", carnumber).First(&carinfo).Error
	return carinfo, err
}
