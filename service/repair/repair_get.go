package repair

import (
	"ProManageSystem/model/repair"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type RepairGetService struct {
	Id          int `form:"id"`
	Ownername   string
	Managername string
}

func (service *RepairGetService) RepairGetOwner() serializer.Response {
	replist, err := repair.GetRepairbyOname(service.Ownername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Success, replist)
}

func (service *RepairGetService) RepairGetManager() serializer.Response {
	replist, err := repair.GetRepairbyMname(service.Managername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Success, replist)
}

func (service *RepairGetService) RepairGet() serializer.Response {
	rep, err := repair.GetRepairbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Success, rep)
}

func (service *RepairGetService) RepairGetTotal() serializer.Response {
	count, err := repair.GetRepairTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *RepairGetService) RepairGetPage(pageindex, pagesize int) serializer.Response {
	replist, err := repair.GetRepairPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	return serializer.GetResponse(serializer.Success, replist)
}
