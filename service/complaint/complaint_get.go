package complaint

import (
	"ProManageSystem/model/complaint"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ComplaintGetService struct {
	Id int `form:"id"`
}

func (service *ComplaintGetService) ComplaintGetMe(ownername string) serializer.Response {
	data, err := complaint.GetComplaintbyOname(ownername)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	return serializer.GetResponse(serializer.Sucess, data)
}
func (service *ComplaintGetService) ComplaintGet() serializer.Response {
	data, err := complaint.GetComplpaintbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	return serializer.GetResponse(serializer.Sucess, data)
}

func (service *ComplaintGetService) ComplaintGetTotal() serializer.Response {
	count, err := complaint.GetComplaintTotal()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Sucess, gin.H{"count": count})
}

func (service *ComplaintGetService) ComplaintGetPage(pageindex, pagesize int) serializer.Response {
	comlist, err := complaint.GetComplpaintPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Sucess, comlist)
}
