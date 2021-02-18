package repair

import (
	"ProManageSystem/model/repair"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type RepairGetService struct {
	Id int `form:"id"`
}

func (service *RepairGetService) RepairGetMe(ownername string) serializer.Response {
	replist, err := repair.GetRepairbyOname(ownername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Sucess, replist)
}

func (service *RepairGetService) RepairGet() serializer.Response {
	rep, err := repair.GetRepairbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Sucess, rep)
}

func (service *RepairGetService) RepairGetTotal() serializer.Response {
	count, err := repair.GetRepairTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Sucess, gin.H{"count": count})
}

func (service *RepairGetService) RepairGetPage(pageindex, pagesize int) serializer.Response {
	replist, err := repair.GetRepairPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Sucess, replist)
}
