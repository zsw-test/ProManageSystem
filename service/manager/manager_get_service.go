package manager

import (
	"ProManageSystem/model/manager"
	"ProManageSystem/serializer"
)

type ManagerGetService struct {
	Id int `form:"id"`
}

func (service *ManagerGetService) ManagerGet() serializer.Response {
	manager, err := manager.GetManagerbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.NotExistUser,
			Result: serializer.GetResult(serializer.NotExistUser),
		}
	} else {
		return serializer.Response{
			Code:   serializer.Sucess,
			Result: serializer.GetResult(serializer.Sucess),
			Data:   manager,
		}
	}
}
