package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ParkInfoTotalService struct {
}

func (serivice *ParkInfoTotalService) ParkInfoTotal(keyword string) serializer.Response {
	count, err := parkmodel.GetParkInfoTotal(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}
