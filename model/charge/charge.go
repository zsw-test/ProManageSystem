package charge

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

type Charge struct {
	gorm.Model
	//根据唯一索引检索到charge收费表  每一个房对应一个表
	Houseid int `gorm:"unique_index:houseid"`
	//还有多少水
	Water float64
	//还有多少度电
	Electric float64
	//还有多少燃气
	Gas float64
	//我的物业费还有多少
	Property float64
}

//增加
func (c *Charge) Create() error {
	err := DB.Mysqldb.Create(c).Error
	return err
}

//保存
func (c *Charge) Save() error {
	err := DB.Mysqldb.Save(c).Error
	return err
}

//删除
func (c *Charge) Delete() error {
	//硬删除永久删除
	err := DB.Mysqldb.Unscoped().Delete(c).Error
	return err
}

//查找
func GetChargebyid(id int) (*Charge, error) {
	var charge = &Charge{}
	err := DB.Mysqldb.Where("id = ?", id).First(&charge).Error
	return charge, err
}

//查找
func GetChargebyhouseid(houseid int) (*Charge, error) {
	var charge = &Charge{}
	err := DB.Mysqldb.Where("houseid = ?", houseid).First(&charge).Error
	return charge, err
}

//获取页面
func GetChargePage(pageindex, pagesize int, keyword string) ([]Charge, error) {
	ChargeList := []Charge{}
	err := DB.Mysqldb.Where("houseid LIKE ?", "%"+keyword+"%").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&ChargeList).Error
	return ChargeList, err
}

func GetChargeTotal(keyword string) (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&Charge{}).Where("houseid LIKE ?", "%"+keyword+"%").Count(&count).Error
	return count, err
}
