package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"
)

type HouseSaveService struct {
	Id           int
	Prorityright int
	HouseType    string
	Ownerid      int
}

func (service *HouseSaveService) HouseSave() serializer.Response {
	house, err := house.GetHousebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorHouseGet)
	}
	house.Prorityright = service.Prorityright
	house.HouseType = service.HouseType
	house.Ownerid = service.Ownerid
	err = house.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorHouseSave)
	}
	return serializer.GetResponse(serializer.Sucess)
}
func (service *HouseSaveService) HouseBuy() serializer.Response {
	house, err := house.GetHousebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorHouseGet)
	}
	house.Ownerid = service.Ownerid
	err = house.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorHouseSave)
	}
	return serializer.GetResponse(serializer.Sucess)
}
