package complaint

import (
	"ProManageSystem/model/complaint"
	"ProManageSystem/serializer"
)

type ComplaintDeleteService struct {
	Id int `form:"id"`
}

func (service *ComplaintDeleteService) ComplaintDelete() serializer.Response {
	complaint, err := complaint.GetComplaintbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	err = complaint.Delete()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintDelete)
	}
	return serializer.GetResponse(serializer.Success)
}
