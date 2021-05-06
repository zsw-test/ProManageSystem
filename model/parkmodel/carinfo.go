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
	PersonalCar string = "私家车"
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

func GetCarinfoPage(pageindex, pagesize int, keyword string) ([]CarInfo, error) {
	CarinfoList := []CarInfo{}
	err := DB.Mysqldb.Where("carnumber LIKE ?", "%"+keyword+"%").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&CarinfoList).Error
	return CarinfoList, err
}

func GetCarinfoTotal(keyword string) (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&CarInfo{}).Where("carnumber LIKE ?", "%"+keyword+"%").Count(&count).Error
	return count, err
}

func GetCarInfobyCarnumber(carnumber string) (*CarInfo, error) {
	var carinfo = &CarInfo{}
	err := DB.Mysqldb.Where("carnumber = ?", carnumber).First(&carinfo).Error
	return carinfo, err
}
