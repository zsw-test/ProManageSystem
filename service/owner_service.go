package service

import (
	"ProManageSystem/model"
)

type OwnerService struct {
	Oname      string `form:"oname"`
	Osex       string `form:"osex"`
	Otelephone int    `form:"otelephone"`
	Oroom      string `form:"oroom"`
	Username   string `form:"username"`
	Upassword  string `form:"password"`
}

func (service *OwnerService) OwnerCreate(owner *model.Owner) error {
	err := owner.Create()
	return err
}

func (service *OwnerService) OwnerSave(owner *model.Owner) error {
	err := owner.Save()
	return err
}
func (service *OwnerService) OwnerGet(username string) (*model.Owner, error) {
	owner, err := model.GetOwner(username)
	return owner, err
}

func (service *OwnerService) OwnerDelete(username string) error {
	owner, err := model.GetOwner(username)
	if err != nil {
		return err
	}
	err = owner.Delete()
	return err
}
