package manager

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type ManagerDeleteService struct {
	Id int `form:"id"`
}

func (service *ManagerDeleteService) ManagerDelete() seralizer.Response {
	manager, err := model.GetManagerbyid(service.Id)
	if err != nil {
		return seralizer.Response{
			Code:   seralizer.NotExistUser,
			Result: "用户不存在",
		}
	}
	err = manager.Delete()
	if err != nil {
		return seralizer.Response{
			Code:   seralizer.ErrorDelete,
			Result: "删除失败",
		}
	}
	return seralizer.Response{
		Code:   seralizer.Sucess,
		Result: "删除成功",
	}
}
