package model

import (
	"github.com/jinzhu/gorm"
)

//停车信息（收费，出入时间）
type ParkInfo struct {
	gorm.Model
	//停车位id
	ParkId int
	//车辆号码
	CarNumber string
	//每小时费用
	Fee int
}
