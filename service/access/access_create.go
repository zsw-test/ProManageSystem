package access

import (
	"ProManageSystem/model/accessrecord"
	"ProManageSystem/serializer"
)

type AccessCreateService struct {
	Name string
	Way  string
}

func (service *AccessCreateService) AccessCreate() serializer.Response {
	exp := accessrecord.Access{
		Name: service.Name,
		Way:  service.Way,
	}
	err := exp.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorRegister,
			Result: serializer.GetResult(serializer.ErrorRegister),
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
		Data:   map[string]interface{}{"id": exp.ID},
	}
}
