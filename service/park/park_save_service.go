package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkSaveService struct {
	Ownerid  int    `form:"ownerid"`
	Status   int8   `form:"status"`
	Location string `form:"location"`
}

func (service *ParkSaveService) ParkSave() serializer.Response {
	park := parkmodel.Park{
		Ownerid:  service.Ownerid,
		Status:   service.Status,
		Location: service.Location,
	}
	err := park.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}
	return serializer.GetResponse(serializer.Sucess)
}
