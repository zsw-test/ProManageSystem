package owner

import (
	"ProManageSystem/DB"

	"github.com/jinzhu/gorm"
)

//业主结构体
type Owner struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);unique_index:username"` //添加唯一索引，防止用户名相同
	Password  string `gorm:"type:varchar(20);"`
	Address   string `gorm:"type:varchar(20);"`
	Telephone int
}

func (user *Owner) Create() error {
	err := DB.Mysqldb.Create(user).Error
	return err
}

func (user *Owner) Save() error {
	err := DB.Mysqldb.Save(user).Error
	return err
}

func (o *Owner) Delete() error {
	err := DB.Mysqldb.Unscoped().Delete(o).Error
	return err
}

func GetOwnerbyname(username string) (*Owner, error) {
	var owner = &Owner{}
	err := DB.Mysqldb.Where("username = ?", username).First(&owner).Error
	return owner, err
}
func GetOwnerbyid(id int) (*Owner, error) {
	var owner = &Owner{}
	err := DB.Mysqldb.Where("id = ?", id).First(&owner).Error
	return owner, err
}

func GetOwnerPage(pageindex, pagesize int) ([]Owner, error) {
	OwnerList := []Owner{}
	err := DB.Mysqldb.Offset((pageindex - 1) * pagesize).Limit(pagesize).Find(&OwnerList).Error
	return OwnerList, err
}

//检查业主权限
func CheckOwnerAuth(username, password string) (bool, error) {
	var auth Owner
	err := DB.Mysqldb.Select("id").Where(Owner{Username: username, Password: password}).First(&auth).Error
	if auth.ID > 0 {
		return true, err
	}

	return false, err
}
