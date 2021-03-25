package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ResidnetGetService struct {
	Houseid    int
	ResidentId int
}

func (service *ResidnetGetService) ResidentGetByid() serializer.Response {
	data, err := house.GetResidentbyid(service.ResidentId)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetResident)
	}
	return serializer.GetResponse(serializer.Success, data)
}

func (service *ResidnetGetService) ResidentGetByhouseidPage(pageindex, pagesize int) serializer.Response {
	data, err := house.GetResidentsbyhouseidPage(service.Houseid, pagesize, pageindex)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetResident)
	}
	return serializer.GetResponse(serializer.Success, data)
}
func (service *ResidnetGetService) ResidentGetByhouseidTotal() serializer.Response {
	count, err := house.GetResidentsbyhouseidTotal(service.Houseid)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetResident)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *ResidnetGetService) ResidentGetPage(pageindex, pagesize int) serializer.Response {
	data, err := house.GetResidentPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetResident)
	}
	return serializer.GetResponse(serializer.Success, data)
}
func (service *ResidnetGetService) ResidentGetTotal() serializer.Response {
	count, err := house.GetResidentTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetResident)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}
