package complaint

import (
	"ProManageSystem/model/complaint"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ComplaintGetService struct {
	Id          int `form:"id"`
	Ownername   string
	Managername string
}

func (service *ComplaintGetService) ComplaintGetMe(ownername string) serializer.Response {
	data, err := complaint.GetComplaintbyOname(ownername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	return serializer.GetResponse(serializer.Success, data)
}
func (service *ComplaintGetService) ComplaintGet() serializer.Response {
	data, err := complaint.GetComplaintbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	return serializer.GetResponse(serializer.Success, data)
}

func (service *ComplaintGetService) ComplaintGetTotal() serializer.Response {
	count, err := complaint.GetComplaintTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *ComplaintGetService) ComplaintGetOwner() serializer.Response {
	replist, err := complaint.GetComplaintbyOname(service.Ownername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	return serializer.GetResponse(serializer.Success, replist)
}

func (service *ComplaintGetService) ComplaintGetManager() serializer.Response {
	replist, err := complaint.GetComplaintbyMname(service.Managername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	return serializer.GetResponse(serializer.Success, replist)
}

func (service *ComplaintGetService) ComplaintGetPage(pageindex, pagesize int) serializer.Response {
	replist, err := complaint.GetComplaintPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	return serializer.GetResponse(serializer.Success, replist)
}
