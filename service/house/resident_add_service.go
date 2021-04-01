package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"
)

type ResidentAddService struct {
	Name    string
	Age     int
	Sex     string
	Work    string
	IdCard  string
	HouseId int
	Address string
}

func (service *ResidentAddService) ResidentAdd() serializer.Response {
	resident := &house.Resident{
		Name:    service.Name,
		Age:     service.Age,
		Sex:     service.Sex,
		Work:    service.Work,
		IdCard:  service.IdCard,
		HouseId: service.HouseId,
		Address: service.Address,
	}

	err := resident.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorAddResident)
	}
	return serializer.GetResponse(serializer.Success)
}
