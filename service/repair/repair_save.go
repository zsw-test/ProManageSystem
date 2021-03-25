package repair

import (
	"ProManageSystem/model/repair"
	"ProManageSystem/serializer"
)

type RepairSaveService struct {
	Id          int    `form:"id"`
	Ownername   string `form:"ownername"`
	Address     string `form:"owneraddress"`
	Reason      string `form:"reason"`
	Status      string `form:"status"`
	Resolve     bool   `form:"resolve"`
	Pics        string
	Managername string
}

func (service *RepairSaveService) RepairSave() serializer.Response {
	rep, err := repair.GetRepairbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	rep.Ownername = service.Ownername
	rep.Address = service.Address
	rep.Reason = service.Reason
	rep.Status = service.Status
	rep.Resolve = service.Resolve
	rep.Pics = service.Pics

	err = rep.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairSave)
	}
	return serializer.GetResponse(serializer.Success)
}

func (service *RepairSaveService) RepairDispatch() serializer.Response {
	rep, err := repair.GetRepairbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	rep.Managername = service.Managername
	rep.Status = "正在处理中"
	err = rep.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairSave)
	}
	return serializer.GetResponse(serializer.Success)
}
func (service *RepairSaveService) RepairResolve() serializer.Response {
	rep, err := repair.GetRepairbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	rep.Status = "已解决"
	rep.Resolve = true
	err = rep.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairSave)
	}
	return serializer.GetResponse(serializer.Success)
}
