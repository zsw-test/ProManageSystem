package carinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type CarinfoGetService struct {
}

func (service *CarinfoGetService) CarinfoGet(carnumber string) serializer.Response {
	data, err := parkmodel.GetCarInfobyCarnumber(carnumber)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorCarinfoGet)
	}
	return serializer.GetResponse(serializer.Success, data)
}

func (service *CarinfoGetService) CarinfoGetTotal(keyword string) serializer.Response {
	count, err := parkmodel.GetCarinfoTotal(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorCarinfoGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *CarinfoGetService) CarinfoGetPage(pageindex, pagesize int, keyword string) serializer.Response {
	comlist, err := parkmodel.GetCarinfoPage(pageindex, pagesize, keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorCarinfoGet)
	}
	return serializer.GetResponse(serializer.Success, comlist)
}
