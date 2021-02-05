package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ParkInfoTotalService struct {
}

func (serivice *ParkInfoTotalService) ParkInfoTotal() serializer.Response {
	count, err := parkmodel.GetParkInfoTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Sucess, gin.H{"count": count})
}
