package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkGetPageService struct {
}

func (service *ParkGetPageService) ParkGetPage(pageindex, pagesize int, keyword string) serializer.Response {
	parklist, err := parkmodel.GetParkPage(pageindex, pagesize, keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, parklist)
}
