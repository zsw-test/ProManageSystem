package manager

import (
	"ProManageSystem/model/manager"
	"ProManageSystem/serializer"
)

type ManagerDeleteService struct {
	Id int `form:"id"`
}

func (service *ManagerDeleteService) ManagerDelete() serializer.Response {
	manager, err := manager.GetManagerbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.NotExistUser,
			Result: "用户不存在",
		}
	}
	err = manager.Delete()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorDelete,
			Result: "删除失败",
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: "删除成功",
	}
}
