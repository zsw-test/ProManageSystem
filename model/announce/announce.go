package announce

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

// 公告
type Announce struct {
	gorm.Model
	// 标题
	Title string
	// 内容
	Text string
	// 发布人
	Managername string
}

//增加
func (c *Announce) Create() error {
	err := DB.Mysqldb.Create(c).Error
	return err
}

//保存
func (c *Announce) Save() error {
	err := DB.Mysqldb.Save(c).Error
	return err
}

//删除
func (c *Announce) Delete() error {
	//硬删除永久删除
	err := DB.Mysqldb.Unscoped().Delete(c).Error
	return err
}

//查找
func GetAnnouncebyid(id int) (*Announce, error) {
	var Announce = &Announce{}
	err := DB.Mysqldb.Where("id = ?", id).First(&Announce).Error
	return Announce, err
}

// 查找所有
func GetAnnounceAll() ([]Announce, error) {
	AnnounceList := []Announce{}
	err := DB.Mysqldb.Order("updated_at desc").Find(&AnnounceList).Error
	return AnnounceList, err
}
