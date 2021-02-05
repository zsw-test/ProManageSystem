package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
	"ProManageSystem/util"
	"time"
)

type ParkInfoGetPageService struct {
}

func (service *ParkInfoGetPageService) ParkInfoGetPage(pageindex, pagesize int) serializer.Response {
	parkinfolist, err := parkmodel.GetParkInfoPage(pageindex, pagesize)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	//模糊计算费用
	for _, parkinfo := range parkinfolist {
		parkinfo.Fee = util.Calcutefee(parkinfo.CreatedAt, time.Now())
		parkinfo.Save()
		if err != nil {
			return serializer.GetResponse(serializer.ErrorUpdateParkFee)
		}
	}
	return serializer.GetResponse(serializer.Sucess, parkinfolist)
}
