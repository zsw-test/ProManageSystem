package repair

import (
	"ProManageSystem/model/repair"
	"ProManageSystem/serializer"
)

type RepairDeleteService struct {
	Id int `form:"id"`
}

func (service *RepairDeleteService) RepairDelete() serializer.Response {
	rep, err := repair.GetRepairbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	err = rep.Delete()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairDelete)
	}
	return serializer.GetResponse(serializer.Sucess)
}
