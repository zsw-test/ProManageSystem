package manager

import (
	"ProManageSystem/model/manager"
	"ProManageSystem/serializer"
)

type ManagerRegisterService struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Depart    string `from:"depart"`
	Telephone string `form:"telephone"`
}

func (service *ManagerRegisterService) ManagerRegister() serializer.Response {
	manager := manager.Manager{
		Username:  service.Username,
		Password:  service.Password,
		Depart:    service.Depart,
		Telephone: service.Telephone,
	}
	err := manager.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ExistUser,
			Result: "用户已经存在",
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: "创建成功",
	}
}
