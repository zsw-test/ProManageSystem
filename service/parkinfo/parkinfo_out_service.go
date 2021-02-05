package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
	"ProManageSystem/util"
	"time"
)

type ParkInfoOutService struct {
	CarNumber string `from:"carnumber"`
}

//出库操作  等到付款完毕 计算费用
func (service *ParkInfoOutService) ParkInfoOut() serializer.Response {
	parkinfo, err := parkmodel.GetParkInfobyCarnumber(service.CarNumber)
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorNotExistRecord,
			Result: serializer.GetResult(serializer.ErrorNotExistRecord),
		}
	}
	//更新计算费用  精确计算费用
	parkinfo.Fee = util.Calcutefee(parkinfo.CreatedAt, time.Now())
	err = parkinfo.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}

	//获取车辆信息
	carinfo, err := parkmodel.GetCarInfobyCarnumber(parkinfo.CarNumber)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}

	if carinfo.CarType != parkmodel.TempoaryCar {
		if carinfo.ExpireTime.After(time.Now()) {
			parkinfo.Fee = 0
			err = parkinfo.Save()
			if err != nil {
				return serializer.GetResponse(serializer.ErrorSave)
			}
			return serializer.GetResponse(serializer.Sucess, parkinfo)
		} else {
			var lasttime time.Time
			if carinfo.ExpireTime.After(parkinfo.CreatedAt) {
				//如果过期时间在后面 那么用过期时间来计算
				lasttime = carinfo.ExpireTime
			} else {
				lasttime = parkinfo.CreatedAt
			}
			parkinfo.Fee = util.Calcutefee(lasttime, time.Now())
			err = parkinfo.Save()
			if err != nil {
				return serializer.GetResponse(serializer.ErrorSave)
			}
			return serializer.GetResponse(serializer.Sucess, parkinfo)
		}
	}

	return serializer.GetResponse(serializer.Sucess, parkinfo)
}
