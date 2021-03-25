package charge

import (
	"ProManageSystem/model/charge"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ChargeRecordGetService struct {
}

func (service *ChargeRecordGetService) ChargeRecordGet(houseid int) serializer.Response {
	chargerecord, err := charge.GetChargesRecordsbyhouseid(houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeRecordGet)
	}
	return serializer.GetResponse(serializer.Success, chargerecord)
}

func (service *ChargeRecordGetService) ChargeRecordGetTotal() serializer.Response {
	count, err := charge.GetChargeRecordTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *ChargeRecordGetService) ChargeRecordGetPage(pageindex, pagesize int) serializer.Response {
	comlist, err := charge.GetChargeRecordPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, comlist)
}
