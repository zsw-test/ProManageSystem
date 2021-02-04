package parkmodel

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

const (
	//占有
	Occupy = 1
	//空
	Empty = 0
	//每小时费用
	HourFee = 5
)

//车位模型
type Park struct {
	gorm.Model
	//所有者
	Ownerid int
	//当前车位的状态
	Status int8
	//地点
	Location string `grom:"varchar(20);unique_index:location"`
}

func (park *Park) Create() error {
	err := DB.Mysqldb.Create(park).Error
	return err
}

func (park *Park) Save() error {
	err := DB.Mysqldb.Save(park).Error
	return err
}

func (park *Park) Delete() error {
	err := DB.Mysqldb.Delete(park).Error
	return err
}

func GetParkbyid(id int) (*Park, error) {
	var park = &Park{}
	err := DB.Mysqldb.Where("id = ?", id).Find(&park).Error
	return park, err
}

func GetParkPage(pageindex, pagesize int) ([]Park, error) {
	parkList := []Park{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&parkList).Error
	return parkList, err
}
