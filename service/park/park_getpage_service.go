package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkGetPageService struct {
}

func (service *ParkGetPageService) ParkGetPage(pageindex, pagesize int) serializer.Response {
	parklist, err := parkmodel.GetParkPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Sucess, parklist)
}
