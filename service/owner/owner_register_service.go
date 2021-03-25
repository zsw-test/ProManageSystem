package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
)

type OwnerRegisterService struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Telephone string `form:"telephone"`
	Houseid   int    `form:"houseid"`
}

func (service *OwnerRegisterService) OwnerRegister() serializer.Response {
	owner := owner.Owner{
		Username:  service.Username,
		Password:  service.Password,
		Telephone: service.Telephone,
		Houseid:   service.Houseid,
	}
	err := owner.Create()
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
