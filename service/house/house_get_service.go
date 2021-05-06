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

func (service *HouseGetService) HouseGetTotal(keyword string) serializer.Response {
	count, err := house.GetHouseTotal(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *HouseGetService) HouseGetPage(pageindex, pagesize int, keyword string) serializer.Response {
	houselist, err := house.GetHousePage(pageindex, pagesize, keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, houselist)
}
