package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkInfoDeleteService struct {
	Id int `from:"id"`
}

func (service *ParkInfoDeleteService) ParkInfoDelete() serializer.Response {
	parkinfo, err := parkmodel.GetParkInfobyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	parkinfo.Delete()
	return serializer.GetResponse(serializer.Sucess)
}
