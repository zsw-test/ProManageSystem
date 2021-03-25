package expressage

import (
	"ProManageSystem/model/expressage"
	"ProManageSystem/serializer"
)

type ExpressageDeleteService struct {
	Id int `form:"id"`
}

func (service *ExpressageDeleteService) ExpressageDelete() serializer.Response {
	complaint, err := expressage.GetExpressagebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	err = complaint.Delete()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageDelete)
	}
	return serializer.GetResponse(serializer.Success)
}
