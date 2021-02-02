package manager

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type ManagerRegisterService struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Depart    string `from:"depart"`
	Telephone int    `form:"telephone"`
}

func (service *ManagerRegisterService) ManagerRegister() seralizer.Response {
	manager := model.Manager{
		Username:  service.Username,
		Password:  service.Password,
		Depart:    service.Depart,
		Telephone: service.Telephone,
	}
	err := manager.Create()
	if err != nil {
		return seralizer.Response{
			Code:   seralizer.ExistUser,
			Result: "用户已经存在",
		}
	}
	return seralizer.Response{
		Code:   seralizer.Sucess,
		Result: "创建成功",
	}
}
