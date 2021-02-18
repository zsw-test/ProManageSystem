package complaint

import (
	"ProManageSystem/model/complaint"
	"ProManageSystem/serializer"
)

type ComplaintCreateService struct {
	Ownername   string `form:"ownername"`
	Managername string `form:"managername"`
	Depart      string `form:"depart"`
	Reason      string `form:"reason"`
}

func (service *ComplaintCreateService) ComplaintCreate() serializer.Response {
	com := &complaint.Complaint{
		Ownername:   service.Ownername,
		Managername: service.Managername,
		Depart:      service.Depart,
		Reason:      service.Reason,
		Status:      "提交",
		Resolve:     false,
	}
	err := com.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorComplaintCreate,
			Result: serializer.GetResult(serializer.ErrorComplaintCreate),
		}
	}
	return serializer.Response{
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
		Data:   map[string]interface{}{"id": com.ID},
	}
}
