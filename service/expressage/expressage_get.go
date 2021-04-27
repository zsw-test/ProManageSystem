package expressage

import (
	"ProManageSystem/model/expressage"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ExpressageGetService struct {
	Id        int `form:"id"`
	Ownername string
}

func (service *ExpressageGetService) ExpressageGetMe() serializer.Response {
	explist, err := expressage.GetExpressagebyOname(service.Ownername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Success, explist)
}

func (service *ExpressageGetService) ExpressageGetTelephone(telephone string) serializer.Response {
	explist, err := expressage.GetExpressagebyTelephone(telephone)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Success, explist)
}

func (service *ExpressageGetService) ExpressageGet() serializer.Response {
	exp, err := expressage.GetExpressagebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Success, exp)
}

func (service *ExpressageGetService) ExpressageGetTotal() serializer.Response {
	count, err := expressage.GetExpressageTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *ExpressageGetService) ExpressageGetPage(pageindex, pagesize int) serializer.Response {
	explist, err := expressage.GetExpressagePage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Success, explist)
}
