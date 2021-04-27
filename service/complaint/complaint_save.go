package complaint

import (
	"ProManageSystem/model/complaint"
	"ProManageSystem/serializer"
)

type ComplaintSaveService struct {
	Id          int    `form:"id"`
	Ownername   string `form:"ownername"`
	Managername string `form:"managername"`
	Depart      string `form:"depart"`
	Reason      string `form:"reason"`
	Status      string `form:"status"`
	Resolve     bool   `form:"resolve"`
}

func (service *ComplaintSaveService) ComplaintSave() serializer.Response {
	com, err := complaint.GetComplaintbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintGet)
	}
	com.Ownername = service.Ownername
	com.Managername = service.Managername
	com.Depart = service.Depart
	com.Reason = service.Reason
	com.Status = service.Status
	com.Resolve = service.Resolve
	err = com.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorComplaintSave)
	}
	return serializer.GetResponse(serializer.Success)
}
func (service *ComplaintSaveService) ComplaintDispatch() serializer.Response {
	rep, err := complaint.GetComplaintbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	rep.Managername = service.Managername
	rep.Status = "正在处理中"
	err = rep.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairSave)
	}
	return serializer.GetResponse(serializer.Success)
}
func (service *ComplaintSaveService) ComplaintResolve() serializer.Response {
	rep, err := complaint.GetComplaintbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairGet)
	}
	rep.Status = "已解决"
	rep.Resolve = true
	err = rep.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorRepairSave)
	}
	return serializer.GetResponse(serializer.Success)
}
