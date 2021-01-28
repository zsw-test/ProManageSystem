package model

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

type Owner struct {
	Oname      string
	Osex       string `gorm:"type:varchar(4)`
	Oroom      string `gorm:"type:varchar(30);not null"`
	Otelephone int
	User
}

func (user *Owner) Create() error {
	err := DB.Mysqldb.Create(user).Error
	return err
}
func (o *Owner) AfterCreate(tx *gorm.DB) (err error) {
	user := User{
		Username:  o.Username,
		Upassword: o.Upassword,
		Utype:     o.Utype,
	}
	err = user.Create()
	if err != nil {
		return err
	}
	return
}

func (user *Owner) Save() error {
	err := DB.Mysqldb.Save(user).Error
	return err
}

func (o *Owner) AfterSave(tx *gorm.DB) (err error) {
	user := User{
		Username:  o.Username,
		Upassword: o.Upassword,
		Utype:     o.Utype,
	}
	err = user.Save()
	if err != nil {
		return err
	}
	return
}

func (o *Owner) Delete() error {
	err := DB.Mysqldb.Delete(o).Error
	return err
}

func (o *Owner) AfterDelete(tx *gorm.DB) (err error) {
	user := User{
		Username:  o.Username,
		Upassword: o.Upassword,
		Utype:     o.Utype,
	}
	err = user.Delete()
	if err != nil {
		return err
	}
	return
}

func GetOwner(username string) (*Owner, error) {
	var owner = &Owner{}
	err := DB.Mysqldb.Where("username = ?", username).Find(&owner).Error
	return owner, err
}
