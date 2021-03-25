package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkSaveService struct {
	Id       int
	Ownerid  int    `form:"ownerid"`
	Status   int8   `form:"status"`
	Location string `form:"location"`
}

func (service *ParkSaveService) ParkSave() serializer.Response {
	park, err := parkmodel.GetParkbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	park.Ownerid = service.Ownerid
	park.Status = service.Status
	park.Location = service.Location
	err = park.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}
	return serializer.GetResponse(serializer.Success)
}
