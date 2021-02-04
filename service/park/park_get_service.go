package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkGetService struct {
	Id int `from:"id"`
}

func (service *ParkGetService) ParkGet() serializer.Response {
	park, err := parkmodel.GetParkbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorGetPark,
			Result: serializer.GetResult(serializer.ErrorGetPark),
		}
	}
	return serializer.Response{
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
		Data:   park,
	}
}
