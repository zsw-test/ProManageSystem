package model

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

type Manager struct {
	Mname string
	Msex  string `grom:"type:varchar(4)"`
	Mno   int
	User
}

func (m *Manager) Create() error {
	err := DB.Mysqldb.Create(m).Error
	return err
}
func (m *Manager) AfterCreate(tx *gorm.DB) (err error) {
	user := User{
		Username:  m.Username,
		Upassword: m.Upassword,
		Utype:     m.Utype,
	}
	err = user.Create()
	if err != nil {
		return err
	}
	return
}
func (manager *Manager) Save() error {
	err := DB.Mysqldb.Save(manager).Error
	return err
}
func (m *Manager) AfterSave(tx *gorm.DB) (err error) {
	user := User{
		Username:  m.Username,
		Upassword: m.Upassword,
		Utype:     m.Utype,
	}
	err = user.Save()
	if err != nil {
		return err
	}
	return
}

func (manager *Manager) Delete() error {
	err := DB.Mysqldb.Delete(manager).Error
	return err
}
func (m *Manager) AfterDelete(tx *gorm.DB) (err error) {
	user := User{
		Username:  m.Username,
		Upassword: m.Upassword,
		Utype:     m.Utype,
	}
	err = user.Delete()
	if err != nil {
		return err
	}
	return
}

func GetManager(username string) (*Manager, error) {
	var manager = &Manager{}
	err := DB.Mysqldb.Where("username = ?", username).Find(&manager).Error
	return manager, err
}
