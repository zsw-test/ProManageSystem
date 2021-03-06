package manager

import (
	"ProManageSystem/model/manager"
	"ProManageSystem/serializer"
)

type ManagerSaveService struct {
	Id        int    `form:"id"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	Depart    string `from:"depart"`
	Telephone string `form:"telephone"`
	Nickname  string
}

func (service *ManagerSaveService) ManagerSave() serializer.Response {
	manager, err := manager.GetManagerbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.NotExistUser,
			Result: "用户不存在",
		}
	}
	manager.Username = service.Username
	manager.Password = service.Password
	manager.Depart = service.Depart
	manager.Telephone = service.Telephone
	manager.Nickname = service.Nickname
	err = manager.Save()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorSave,
			Result: "保存失败",
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: "保存成功",
	}
}
