package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
	"ProManageSystem/util"
	"time"
)

type ParkInfoGetPageService struct {
}

func (service *ParkInfoGetPageService) ParkInfoGetPage(pageindex, pagesize int, keyword string) serializer.Response {
	parkinfolist, err := parkmodel.GetParkInfoPage(pageindex, pagesize, keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	//模糊计算费用
	for _, parkinfo := range parkinfolist {
		parkinfo.Fee = util.Calcutefee(parkinfo.CreatedAt, time.Now())
		parkinfo.Parktime = time.Since(parkinfo.CreatedAt).String()
		parkinfo.Save()
		if err != nil {
			return serializer.GetResponse(serializer.ErrorUpdateParkFee)
		}
	}
	return serializer.GetResponse(serializer.Success, parkinfolist)
}
