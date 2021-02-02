package owner

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type OwnerRegisterService struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	Telephone int    `form:"telephone"`
	Address   string `form:"address"`
}

func (service *OwnerRegisterService) OwnerRegister() seralizer.Response {
	owner := model.Owner{
		Username:  service.Username,
		Password:  service.Password,
		Telephone: service.Telephone,
		Address:   service.Address,
	}
	err := owner.Create()
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
