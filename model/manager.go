package model

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

type Manager struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);unique_index:username"` //添加唯一索引，防止用户名相同
	Password  string `gorm:"type:varchar(20);"`
	Depart    string `gorm:"type:varchar(20);"`
	Telephone int
}

func (m *Manager) Create() error {
	err := DB.Mysqldb.Create(m).Error
	return err
}

func (manager *Manager) Save() error {
	err := DB.Mysqldb.Save(manager).Error
	return err
}

func (manager *Manager) Delete() error {
	err := DB.Mysqldb.Delete(manager).Error
	return err
}

func GetManagerbyname(username string) (*Manager, error) {
	var manager = &Manager{}
	err := DB.Mysqldb.Where("username = ?", username).Find(&manager).Error
	return manager, err
}
func GetManagerbyid(id int) (*Manager, error) {
	var manager = &Manager{}
	err := DB.Mysqldb.Where("id = ?", id).Find(&manager).Error
	return manager, err
}

func GetManagerPage(pageindex, pagesize int) ([]Manager, error) {
	ManagerList := []Manager{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&ManagerList).Error
	return ManagerList, err
}
