package manager

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type ManagerSaveService struct {
	Id        int    `form:"id"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	Depart    string `from:"depart"`
	Telephone int    `form:"telephone"`
}

func (service *ManagerSaveService) ManagerSave() seralizer.Response {
	manager, err := model.GetManagerbyid(service.Id)
	if err != nil {
		return seralizer.Response{
			Code:   "404",
			Result: "用户不存在",
		}
	}
	manager.Username = service.Username
	manager.Password = service.Password
	manager.Depart = service.Depart
	manager.Telephone = service.Telephone
	err = manager.Save()
	if err != nil {
		return seralizer.Response{
			Code:   "404",
			Result: "保存失败",
		}
	}
	return seralizer.Response{
		Code:   "200",
		Result: "保存成功",
	}
}
