package owner

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type OwnerGetService struct {
	Id int `form:"id"`
}

func (service *OwnerGetService) OwnerGet() seralizer.Response {
	owner, err := model.GetOwnerbyid(service.Id)
	if err != nil {
		return seralizer.Response{
			Code:   "404",
			Result: "不存在用户",
		}
	} else {
		return seralizer.Response{
			Code:   "200",
			Result: "查询成功",
			Data:   owner,
		}
	}
}
