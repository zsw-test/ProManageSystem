package charge

import (
	"ProManageSystem/model/charge"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ChargeGetService struct {
}

func (service *ChargeGetService) ChargeGetWater(houseid int) serializer.Response {
	charge, err := charge.GetChargebyhouseid(houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"剩余水量": charge.Water})
}
func (service *ChargeGetService) ChargeGetElectric(houseid int) serializer.Response {
	charge, err := charge.GetChargebyhouseid(houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"剩余电量": charge.Electric})
}
func (service *ChargeGetService) ChargeGetGas(houseid int) serializer.Response {
	charge, err := charge.GetChargebyhouseid(houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"剩余燃气": charge.Gas})
}
func (service *ChargeGetService) ChargeGetProperty(houseid int) serializer.Response {
	charge, err := charge.GetChargebyhouseid(houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"需缴纳物业费": charge.Property})
}
func (service *ChargeGetService) ChargeGet(houseid int) serializer.Response {
	charge, err := charge.GetChargebyhouseid(houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorChargeGet)
	}
	return serializer.GetResponse(serializer.Success, charge)
}

func (service *ChargeGetService) ChargeGetTotal(keyword string) serializer.Response {
	count, err := charge.GetChargeTotal(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *ChargeGetService) ChargeGetPage(pageindex, pagesize int, keyword string) serializer.Response {
	comlist, err := charge.GetChargePage(pageindex, pagesize, keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, comlist)
}
