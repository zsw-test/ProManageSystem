package expressage

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

//快递信息结构体
type Expressage struct {
	gorm.Model
	Ownername       string
	ExpressLocation string `gorm:"column:expressagelocation"`
	Istake          bool
}

//增加
func (expressage *Expressage) Create() error {

	err := DB.Mysqldb.Create(expressage).Error
	return err
}

//保存
func (expressage *Expressage) Save() error {
	err := DB.Mysqldb.Save(expressage).Error
	return err
}

//删除
func (expressage *Expressage) Delete() error {
	//硬删除永久删除
	err := DB.Mysqldb.Unscoped().Delete(expressage).Error
	return err
}

//根据快递号查找
func GetExpressagebyid(id int) (*Expressage, error) {
	var expressage = &Expressage{}
	err := DB.Mysqldb.Where("id = ?", id).First(&expressage).Error
	return expressage, err
}

//根据用户名查找 可能有多个  所以返回一个list
func GetExpressagebyname(Ownername int) ([]Expressage, error) {
	ExpressageList := []Expressage{}
	err := DB.Mysqldb.Where("ownername = ?", Ownername).Find(&ExpressageList).Error
	return ExpressageList, err
}

//获取页面
func GetExpressagePage(pageindex, pagesize int) ([]Expressage, error) {
	ExpressageList := []Expressage{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&ExpressageList).Error
	return ExpressageList, err
}
