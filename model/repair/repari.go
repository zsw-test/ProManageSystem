package repair

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

type Repair struct {
	gorm.Model
	Ownername    string
	Owneraddress string
	Reason       string
	Reslove      bool
}

//增加
func (repair *Repair) Create() error {
	err := DB.Mysqldb.Create(repair).Error
	return err
}

//保存
func (repair *Repair) Save() error {
	err := DB.Mysqldb.Save(repair).Error
	return err
}

//删除
func (repair *Repair) Delete() error {
	//硬删除永久删除
	err := DB.Mysqldb.Unscoped().Delete(repair).Error
	return err
}

//查找
func GetRepairbyid(id int) (*Repair, error) {
	var repair = &Repair{}
	err := DB.Mysqldb.Where("id = ?", id).First(&repair).Error
	return repair, err
}

//获取页面
func GetRepairPage(pageindex, pagesize int) ([]Repair, error) {
	RepairList := []Repair{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&RepairList).Error
	return RepairList, err
}
