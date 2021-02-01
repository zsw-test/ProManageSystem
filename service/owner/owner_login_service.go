package owner

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type OwnerLoginService struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (service *OwnerLoginService) OwnerLogin() seralizer.Response {
	owner, err := model.GetOwnerbyname(service.Username)
	if err != nil {
		return seralizer.Response{
			Code:   "404",
			Result: "不存在用户",
		}
	}
	if service.Password != owner.Password {
		return seralizer.Response{
			Code:   "404",
			Result: "密码错误",
		}
	}
	return seralizer.Response{
		Code:   "200",
		Result: "登陆成功",
		Data:   owner,
	}
}
