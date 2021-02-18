package repair

import (
	"ProManageSystem/model/repair"
	"ProManageSystem/serializer"
)

type RepairSaveService struct {
	Id           int    `form:"id"`
	Ownername    string `form:"ownername"`
	Owneraddress string `form:"owneraddress"`
	Reason       string `form:"reason"`
	Status       string `form:"status"`
	Resolve      bool   `form:"resolve"`
}

func (service *RepairSaveService) RepairSave() serializer.Response {
	rep, err := repair.GetRepairbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	rep.Ownername = service.Ownername
	rep.Owneraddress = service.Owneraddress
	rep.Reason = service.Reason
	rep.Status = service.Status
	rep.Resolve = service.Resolve

	err = rep.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairSave)
	}
	return serializer.GetResponse(serializer.Sucess)
}
