package charge

import (
	"ProManageSystem/model/charge"
	"ProManageSystem/serializer"
)

type ChargeRecordCreateService struct {
	Houseid int `form:"houseid"`
	//缴费类型
	Chargetype string `form:"chargetype"`
	//缴费费用
	Chargefee float64 `form:"chargefee"`
}

func (service *ChargeRecordCreateService) ChargeRecordCreate() serializer.Response {
	chargerecord := &charge.ChargeRecord{
		Houseid:    service.Houseid,
		Chargetype: service.Chargetype,
		Chargefee:  service.Chargefee,
	}
	err := chargerecord.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeRecordCreate)
	}
	return serializer.GetResponse(serializer.Success)
}
