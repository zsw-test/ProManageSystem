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
	Location string `gorm:"type:varchar(20);unique_index:location"`
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
	err := DB.Mysqldb.Unscoped().Delete(park).Error
	return err
}

func GetParkbyid(id int) (*Park, error) {
	var park = &Park{}
	err := DB.Mysqldb.Where("id = ?", id).Find(&park).Error
	return park, err
}

func GetParkPage(pageindex, pagesize int, keyword string) ([]Park, error) {
	ParkList := []Park{}
	err := DB.Mysqldb.Where("location LIKE ?", "%"+keyword+"%").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&ParkList).Error
	return ParkList, err
}

func GetParkTotal(keyword string) (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&Park{}).Where("location LIKE ?", "%"+keyword+"%").Count(&count).Error
	return count, err
}
func GetParkEmpty() (int, error) {
	count := 0
	err := DB.Mysqldb.Where("Status = ?", 0).Model(&Park{}).Count(&count).Error
	return count, err
}
func GetParkOccupy() (int, error) {
	count := 0
	err := DB.Mysqldb.Where("Status = ?", 1).Model(&Park{}).Count(&count).Error
	return count, err
}
func GetParkPersonal() (int, error) {
	count := 0
	err := DB.Mysqldb.Where("Ownerid <> ?", 0).Model(&Park{}).Count(&count).Error
	return count, err
}
func GetParkFree() (int, error) {
	count := 0
	err := DB.Mysqldb.Where("Ownerid = ?", 0).Model(&Park{}).Count(&count).Error
	return count, err
}
func GetParkFreeList() ([]Park, error) {
	parkList := []Park{}
	err := DB.Mysqldb.Where("Ownerid = ?", 0).Find(&parkList).Error
	return parkList, err
}
