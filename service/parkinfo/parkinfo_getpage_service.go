package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkInfoGetPageService struct {
}

func (service *ParkInfoGetPageService) ParkInfoGetPage(pageindex, pagesize int) serializer.Response {
	parkinfolist, err := parkmodel.GetParkInfoPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	for _, parkinfo := range parkinfolist {
		err := parkinfo.UpdateFee()
		if err != nil {
			return serializer.GetResponse(serializer.ErrorUpdateParkFee)
		}
	}
	return serializer.GetResponse(serializer.Sucess, parkinfolist)
}
