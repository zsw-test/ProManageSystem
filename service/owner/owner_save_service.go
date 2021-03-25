package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
)

type OwnerSaveService struct {
	Id        int    `form:"id"`
	Username  string `form:"username"`
	Password  string `form:"password"`
	Telephone string `form:"telephone"`
	Houseid   int    `form:"houseid"`
	Nickname  string
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
	owner.Houseid = service.Houseid
	owner.Telephone = service.Telephone
	owner.Password = service.Password
	owner.Nickname = service.Nickname
	err = owner.Save()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorSave,
			Result: serializer.GetResult(serializer.ErrorSave),
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
	}
}
