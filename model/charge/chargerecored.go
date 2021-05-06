package charge

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

const (
	//水每立方米多少钱
	WaterCountFee float64 = 4
	//燃气每立方米多少钱
	GasCountFee float64 = 2
	//电费每度多少钱
	ElectricCountFee float64 = 0.5
	//物业费每个月100
	PropertyCountFee float64 = 100
)

type ChargeRecord struct {
	gorm.Model
	Houseid int
	//缴费类型
	Chargetype string
	//缴费费用
	Chargefee float64
}

//增加
func (c *ChargeRecord) Create() error {
	err := DB.Mysqldb.Create(c).Error
	return err
}

//保存
func (c *ChargeRecord) Save() error {
	err := DB.Mysqldb.Save(c).Error
	return err
}

//删除
func (c *ChargeRecord) Delete() error {
	//硬删除永久删除
	err := DB.Mysqldb.Unscoped().Delete(c).Error
	return err
}

//查找
func GetChargeRecordbyid(id int) (*ChargeRecord, error) {
	var chargerecord = &ChargeRecord{}
	err := DB.Mysqldb.Where("id = ?", id).First(&chargerecord).Error
	return chargerecord, err
}

//查找
func GetChargesRecordsbyhouseid(houseid int) ([]ChargeRecord, error) {
	ChargeRecordList := []ChargeRecord{}
	err := DB.Mysqldb.Where("houseid = ?", houseid).First(&ChargeRecordList).Error
	return ChargeRecordList, err
}

//获取页面
func GetChargeRecordPage(pageindex, pagesize int, keyword string) ([]ChargeRecord, error) {
	ChargeRecordList := []ChargeRecord{}
	err := DB.Mysqldb.Where("houseid LIKE ?", "%"+keyword+"%").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&ChargeRecordList).Error
	return ChargeRecordList, err
}

func GetChargeRecordTotal(keyword string) (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&ChargeRecord{}).Where("houseid LIKE ?", "%"+keyword+"%").Count(&count).Error
	return count, err
}
