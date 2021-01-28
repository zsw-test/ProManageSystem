package model

import (
	"ProManageSystem/DB"
)

//用户数据模型（包括物业和业主用户）
type User struct {
	Username  string `gorm:"type:varchar(20);primary_key"`
	Upassword string `gorm:"type:varchar(20);not null"`
	Utype     int8
}

func (user *User) Create() error {
	err := DB.Mysqldb.Create(&user).Error
	return err
}
func (user *User) Save() error {
	err := DB.Mysqldb.Save(&user).Error
	return err
}
func (user *User) Delete() error {
	err := DB.Mysqldb.Delete(&user).Error
	return err
}
func GetUser(username string) (*User, error) {
	user := User{}
	result := DB.Mysqldb.Where("username=?", username).First(&user)
	return &user, result.Error
}
