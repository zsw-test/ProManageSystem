package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ParkGetService struct {
	Id int `from:"id"`
}

func (service *ParkGetService) ParkGet() serializer.Response {
	park, err := parkmodel.GetParkbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorGetPark,
			Result: serializer.GetResult(serializer.ErrorGetPark),
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
		Data:   park,
	}
}

func (service *ParkGetService) ParkGetTotal() serializer.Response {
	count, err := parkmodel.GetParkTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *ParkGetService) ParkGetFreeList() serializer.Response {
	parklist, err := parkmodel.GetParkFreeList()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	return serializer.GetResponse(serializer.Success, parklist)
}
