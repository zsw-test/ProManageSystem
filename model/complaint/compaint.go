package complaint

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

//Complpaint  投诉结构体
type Complaint struct {
	gorm.Model
	Ownername   string
	Managername string
	Depart      string
	Reason      string
	//投诉的当前状态
	Status  string
	Resolve bool
}

//增加
func (c *Complaint) Create() error {
	err := DB.Mysqldb.Create(c).Error
	return err
}

//保存
func (c *Complaint) Save() error {
	err := DB.Mysqldb.Save(c).Error
	return err
}

//删除
func (c *Complaint) Delete() error {
	//硬删除永久删除
	err := DB.Mysqldb.Unscoped().Delete(c).Error
	return err
}

//查找
func GetComplaintbyid(id int) (*Complaint, error) {
	var complaint = &Complaint{}
	err := DB.Mysqldb.Where("id = ?", id).First(&complaint).Error
	return complaint, err
}

//获取页面
func GetComplaintPage(pageindex, pagesize int, keyword string) ([]Complaint, error) {
	ComplaintList := []Complaint{}
	err := DB.Mysqldb.Where("ownername LIKE ?", "%"+keyword+"%").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&ComplaintList).Error
	return ComplaintList, err
}

func GetComplaintTotal(keyword string) (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&Complaint{}).Where("ownername LIKE ?", "%"+keyword+"%").Count(&count).Error
	return count, err
}

//根据用户名查找 可能有多个  所以返回一个list
func GetComplaintbyOname(ownername string) ([]Complaint, error) {
	ComplaintList := []Complaint{}
	err := DB.Mysqldb.Where("ownername = ?", ownername).Find(&ComplaintList).Error
	return ComplaintList, err
}

//根据用户名查找 可能有多个  所以返回一个list
func GetComplaintbyMname(managername string) ([]Complaint, error) {
	ComplaintList := []Complaint{}
	err := DB.Mysqldb.Where("managername = ?", managername).Find(&ComplaintList).Error
	return ComplaintList, err
}
