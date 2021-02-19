package house

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

//房屋模型 包括房屋的信息
type House struct {
	gorm.Model
	//楼栋号码
	Building int
	//单元号码
	Unit         int
	Area         float64
	Prorityright int
	HouseType    string
	Ownerid      int
}

func (house *House) Create() error {
	err := DB.Mysqldb.Create(house).Error
	return err
}

func (house *House) Save() error {
	err := DB.Mysqldb.Save(house).Error
	return err
}

func (house *House) Delete() error {
	err := DB.Mysqldb.Delete(house).Error
	return err
}

func GetHousebyid(id int) (*House, error) {
	var house = &House{}
	err := DB.Mysqldb.Where("id = ?", id).Find(&house).Error
	return house, err
}

func GetHousePage(pageindex, pagesize int) ([]House, error) {
	houseList := []House{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&houseList).Error
	return houseList, err
}
func GetHouseTotal() (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&House{}).Count(&count).Error
	return count, err
}
