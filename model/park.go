package model

import (
	"github.com/jinzhu/gorm"
)

//车位模型
type Park struct {
	gorm.Model
	//所有者
	Ownerid int
	//当前车位的状态
	Status int8
	//地点
	Location string
}
