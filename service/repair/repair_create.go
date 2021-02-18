package repair

import (
	"ProManageSystem/model/repair"
	"ProManageSystem/serializer"
)

type RepairCreateService struct {
	Ownername    string `form:"ownername"`
	Owneraddress string `form:"owneraddress"`
	Reason       string `form:"reason"`
	Status       string `form:"status"`
	Resolve      bool   `form:"resolve"`
}

func (service *RepairCreateService) RepairCreate() serializer.Response {
	rep := &repair.Repair{
		Ownername:    service.Ownername,
		Owneraddress: service.Owneraddress,
		Reason:       service.Reason,
		Status:       service.Status,
		Resolve:      service.Resolve,
	}
	err := rep.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorRepairCreate,
			Result: serializer.GetResult(serializer.ErrorRepairCreate),
		}
	}
	return serializer.Response{
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
		Data:   map[string]interface{}{"id": rep.ID},
	}
}
