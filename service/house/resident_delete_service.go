package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"
)

type ResidentDeleteService struct {
	ResidentId int
}

func (service *ResidentDeleteService) ResidentDelete() serializer.Response {
	resident, err := house.GetResidentbyid(service.ResidentId)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetResident)
	}
	err = resident.Delete()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorDelResident)
	}

	return serializer.GetResponse(serializer.Sucess)
}
