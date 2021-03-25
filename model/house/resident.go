package house

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

//住户信息
type Resident struct {
	gorm.Model
	Name    string
	Age     int
	Sex     string
	Work    string
	IdCard  string
	HouseId int
	Address string
}

func (resident *Resident) Create() error {
	err := DB.Mysqldb.Create(resident).Error
	return err
}

func (resident *Resident) Save() error {
	err := DB.Mysqldb.Save(resident).Error
	return err
}

func (resident *Resident) Delete() error {
	err := DB.Mysqldb.Unscoped().Delete(resident).Error
	return err
}

func GetResidentbyid(id int) (*Resident, error) {
	var resident = &Resident{}
	err := DB.Mysqldb.Where("id = ?", id).Find(&resident).Error
	return resident, err
}

//获取该房间的所有住户
func GetResidentsbyhouseidPage(houseid, pagesize, pageindex int) ([]Resident, error) {
	residentList := []Resident{}
	err := DB.Mysqldb.Where("houseid = ?", houseid).Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&residentList).Error
	return residentList, err
}
func GetResidentsbyhouseidTotal(houseid int) (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&Resident{}).Where("houseid = ?", houseid).Count(&count).Error
	return count, err
}

func GetResidentPage(pageindex, pagesize int) ([]Resident, error) {
	residentList := []Resident{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&residentList).Error
	return residentList, err
}
func GetResidentTotal() (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&Resident{}).Count(&count).Error
	return count, err
}
