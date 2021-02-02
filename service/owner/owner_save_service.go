package owner

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type OwnerSaveService struct {
	Id        int    `form:"id"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	Telephone int    `form:"telephone"`
	Address   string `form:"address"`
}

func (service *OwnerSaveService) OwnerSave() seralizer.Response {
	owner, err := model.GetOwnerbyid(service.Id)
	if err != nil {
		return seralizer.Response{
			Code:   seralizer.NotExistUser,
			Result: seralizer.GetResult(seralizer.NotExistUser),
		}
	}
	owner.Username = service.Username
	owner.Address = service.Address
	owner.Telephone = service.Telephone
	owner.Password = service.Password
	err = owner.Save()
	if err != nil {
		return seralizer.Response{
			Code:   seralizer.ErrorSave,
			Result: seralizer.GetResult(seralizer.ErrorSave),
		}
	}
	return seralizer.Response{
		Code:   seralizer.Sucess,
		Result: seralizer.GetResult(seralizer.Sucess),
	}
}
