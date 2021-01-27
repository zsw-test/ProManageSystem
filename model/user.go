package model

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

//用户数据模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null;unique"`
	Password string `gorm:"type:varchar(20);not null"`
	Address  string `gorm:"type:varchar(255);not null"`
}

func (user *User) Create() error {
	err := DB.Mysqldb.Create(user).Error
	return err
}
func (user *User) Save() error {
	err := DB.Mysqldb.Save(user).Error
	return err
}
func (user *User) Delete() error {
	err := DB.Mysqldb.Delete(user).Error
	return err
}
func GetUser(username string) (*User, error) {
	var user User
	result := DB.Mysqldb.Where(User{Username: username}).First(&user)
	return &user, result.Error
}
