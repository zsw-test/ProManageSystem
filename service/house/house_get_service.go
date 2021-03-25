package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type HouseGetService struct {
	Id int `form:"id"`
}

func (service *HouseGetService) HouseGet() serializer.Response {
	data, err := house.GetHousebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorHouseGet)
	}
	return serializer.GetResponse(serializer.Success, data)
}

func (service *HouseGetService) HouseGetTotal() serializer.Response {
	count, err := house.GetHouseTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *HouseGetService) HouseGetPage(pageindex, pagesize int) serializer.Response {
	houselist, err := house.GetHousePage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, houselist)
}

