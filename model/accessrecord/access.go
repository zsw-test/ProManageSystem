package accessrecord

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

type Access struct {
	gorm.Model
	Name string
	Way  string
}

//增加
func (c *Access) Create() error {
	err := DB.Mysqldb.Create(c).Error
	return err
}

//保存
func (c *Access) Save() error {
	err := DB.Mysqldb.Save(c).Error
	return err
}

//删除
func (c *Access) Delete() error {
	//硬删除永久删除
	err := DB.Mysqldb.Unscoped().Delete(c).Error
	return err
}

//查找
func GetAccessbyid(id int) (*Access, error) {
	var Access = &Access{}
	err := DB.Mysqldb.Where("id = ?", id).First(&Access).Error
	return Access, err
}

// 查找所有
func GetAccessAll(keyword string) ([]Access, error) {
	AccessList := []Access{}
	err := DB.Mysqldb.Where("name LIKE ?", "%"+keyword+"%").Order("updated_at desc").Limit(100).Find(&AccessList).Error
	return AccessList, err
}
