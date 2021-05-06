package manager

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

//物业管理员结构体
type Manager struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);unique_index:username"` //添加唯一索引，防止用户名相同
	Nickname  string `gorm:"type:varchar(20);"`
	Password  string `gorm:"type:varchar(20);"`
	Depart    string `gorm:"type:varchar(20);"`
	Telephone string `gorm:"type:varchar(20);"`
}

func (m *Manager) Create() error {
	err := DB.Mysqldb.Create(m).Error
	return err
}

func (manager *Manager) Save() error {
	err := DB.Mysqldb.Save(manager).Error
	return err
}

func (manager *Manager) Delete() error {
	//硬删除永久删除
	err := DB.Mysqldb.Unscoped().Delete(manager).Error
	return err
}

func GetManagerbyname(username string) (*Manager, error) {
	var manager = &Manager{}
	err := DB.Mysqldb.Where("username = ?", username).First(&manager).Error
	return manager, err
}
func GetManagerbyid(id int) (*Manager, error) {
	var manager = &Manager{}
	err := DB.Mysqldb.Where("id = ?", id).First(&manager).Error
	return manager, err
}

func GetManagerPage(pageindex, pagesize int, keyword string) ([]Manager, error) {
	ManagerList := []Manager{}
	err := DB.Mysqldb.Where("username LIKE ?", "%"+keyword+"%").Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&ManagerList).Error
	return ManagerList, err
}
func GetManagerTotal(keyword string) (int, error) {
	count := 0
	err := DB.Mysqldb.Model(&Manager{}).Where("username LIKE ?", "%"+keyword+"%").Count(&count).Error
	return count, err
}

//检查管理员权限
func CheckManagerAuth(username, password string) (bool, error) {
	var auth Manager
	err := DB.Mysqldb.Select("id").Where(Manager{Username: username, Password: password}).First(&auth).Error
	if auth.ID > 0 {
		return true, err
	}

	return false, err
}
