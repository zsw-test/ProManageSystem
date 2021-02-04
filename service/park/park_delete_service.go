package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkDeleteService struct {
	Id int `form:"id"`
}

func (service *ParkDeleteService) ParkDelete() serializer.Response {
	park, err := parkmodel.GetParkbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorNotExistRecord,
			Result: serializer.GetResult(serializer.ErrorNotExistRecord),
		}
	}
	park.Delete()
	return serializer.Response{
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
	}
}
