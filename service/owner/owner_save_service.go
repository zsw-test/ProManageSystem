package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
)

type OwnerSaveService struct {
	Id        int    `form:"id"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	Telephone int    `form:"telephone"`
	Address   string `form:"address"`
}

func (service *OwnerSaveService) OwnerSave() serializer.Response {
	owner, err := owner.GetOwnerbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.NotExistUser,
			Result: serializer.GetResult(serializer.NotExistUser),
		}
	}
	owner.Username = service.Username
	owner.Address = service.Address
	owner.Telephone = service.Telephone
	owner.Password = service.Password
	err = owner.Save()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorSave,
			Result: serializer.GetResult(serializer.ErrorSave),
		}
	}
	return serializer.Response{
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
	}
}
