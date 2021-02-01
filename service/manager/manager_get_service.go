package manager

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type ManagerGetService struct {
	Id int `form:"id"`
}

func (service *ManagerGetService) ManagerGet() seralizer.Response {
	manager, err := model.GetManagerbyid(service.Id)
	if err != nil {
		return seralizer.Response{
			Code:   "404",
			Result: "不存在用户",
		}
	} else {
		return seralizer.Response{
			Code:   "200",
			Result: "查询成功",
			Data:   manager,
		}
	}
}
