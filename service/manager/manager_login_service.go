package manager

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type ManagerLoginService struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (service *ManagerLoginService) ManagerLogin() seralizer.Response {
	manager, err := model.GetManagerbyname(service.Username)
	if err != nil {
		return seralizer.Response{
			Code:   "404",
			Result: "不存在用户",
		}
	}
	if service.Password != manager.Password {
		return seralizer.Response{
			Code:   "404",
			Result: "密码错误",
		}
	}
	return seralizer.Response{
		Code:   "200",
		Result: "登陆成功",
		Data:   manager,
	}
}
