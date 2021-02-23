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
	return serializer.GetResponse(serializer.Sucess, data)
}

func (service *CarinfoGetService) CarinfoGetTotal() serializer.Response {
	count, err := parkmodel.GetCarinfoTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorCarinfoGet)
	}
	return serializer.GetResponse(serializer.Sucess, gin.H{"count": count})
}

func (service *CarinfoGetService) CarinfoGetPage(pageindex, pagesize int) serializer.Response {
	comlist, err := parkmodel.GetCarInfoPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorCarinfoGet)
	}
	return serializer.GetResponse(serializer.Sucess, comlist)
}
