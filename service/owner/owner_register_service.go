package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
)

type OwnerRegisterService struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Telephone int    `form:"telephone"`
	Address   string `form:"address"`
}

func (service *OwnerRegisterService) OwnerRegister() serializer.Response {
	owner := owner.Owner{
		Username:  service.Username,
		Password:  service.Password,
		Telephone: service.Telephone,
		Address:   service.Address,
	}
	err := owner.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ExistUser,
			Result: "用户已经存在",
		}
	}
	return serializer.Response{
		Code:   serializer.Sucess,
		Result: "创建成功",
	}
}
