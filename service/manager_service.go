package service

import "ProManageSystem/model"

type ManagerService struct {
	Mname    string `form:"mname"`
	Msex     string `form:"msex"`
	Mno      int    `form:"mno"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func (service *ManagerService) ManagerCreate(m *model.Manager) error {
	err := m.Create()
	return err
}

func (service *ManagerService) ManagerSave(m *model.Manager) error {
	err := m.Save()
	return err
}

func (service *ManagerService) ManagerGet(username string) (*model.Manager, error) {
	manager, err := model.GetManager(username)
	return manager, err
}

func (service *ManagerService) ManagerDelete(username string) error {
	manager, err := model.GetManager(username)
	if err != nil {
		return err
	}
	err = manager.Delete()
	return err
}
