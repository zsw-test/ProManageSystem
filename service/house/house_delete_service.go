package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"
)

type HouseDeleteService struct {
	Id int `form:"id"`
}

func (service *HouseDeleteService) HouseDelete() serializer.Response {
	house, err := house.GetHousebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorHouseGet)
	}
	err = house.Delete()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorHouseDelete)
	}
	return serializer.GetResponse(serializer.Success)
}
