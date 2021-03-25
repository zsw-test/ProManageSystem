package complaint

import (
	"ProManageSystem/model/complaint"
	"ProManageSystem/serializer"
)

type ComplaintResolveService struct {
	Id int `form:"id"`
}

func (service *ComplaintResolveService) ComplaintResolve() serializer.Response {
	complaint, err := complaint.GetComplaintbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	complaint.Resolve = true
	complaint.Save()
	return serializer.GetResponse(serializer.Success)
}
