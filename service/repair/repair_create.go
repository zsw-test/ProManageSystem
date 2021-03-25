package repair

import (
	"ProManageSystem/model/repair"
	"ProManageSystem/serializer"
)

type RepairCreateService struct {
	Ownername string `form:"ownername"`
	Address   string `form:"address"`
	Reason    string `form:"reason"`
	Status    string `form:"status"`
	Resolve   bool   `form:"resolve"`
	Pics      string `form:"pics"`
}

func (service *RepairCreateService) RepairCreate() serializer.Response {
	rep := &repair.Repair{
		Ownername: service.Ownername,
		Address:   service.Address,
		Reason:    service.Reason,
		Status:    service.Status,
		Resolve:   service.Resolve,
		Pics:      service.Pics,
	}
	err := rep.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorRepairCreate,
			Result: serializer.GetResult(serializer.ErrorRepairCreate),
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
		Data:   map[string]interface{}{"id": rep.ID},
	}
}
