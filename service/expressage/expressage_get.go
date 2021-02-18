package expressage

import (
	"ProManageSystem/model/expressage"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ExpressageGetService struct {
	Id int `form:"id"`
}

func (service *ExpressageGetService) ExpressageGetMe(ownername string) serializer.Response {
	replist, err := expressage.GetExpressagebyOname(ownername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Sucess, replist)
}

func (service *ExpressageGetService) ExpressageGet() serializer.Response {
	rep, err := expressage.GetExpressagebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Sucess, rep)
}

func (service *ExpressageGetService) ExpressageGetTotal() serializer.Response {
	count, err := expressage.GetExpressageTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Sucess, gin.H{"count": count})
}

func (service *ExpressageGetService) ExpressageGetPage(pageindex, pagesize int) serializer.Response {
	replist, err := expressage.GetExpressagePage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	return serializer.GetResponse(serializer.Sucess, replist)
}
