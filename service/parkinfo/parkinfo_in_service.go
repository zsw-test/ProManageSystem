package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkInfoInService struct {
	//车辆号码
	CarNumber string `form:"carnumber"`
}

//创建一个停车信息  同时检测车辆信息 是否在数据库中 如果不再 则加入   临时车信息
func (service *ParkInfoInService) ParkInfoIn() serializer.Response {

	parkinfo := parkmodel.ParkInfo{
		CarNumber: service.CarNumber,
	}
	err := parkinfo.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorCreateParkinfo)
	}
	carinfo, err := parkmodel.GetCarInfobyCarnumber(service.CarNumber)
	if err != nil {
		//如果车辆信息不存在  那么创建一个临时车信息
		carinfo = &parkmodel.CarInfo{
			CarNumber: service.CarNumber,
			CarType:   parkmodel.TempoaryCar,
		}
		err := carinfo.Create()
		if err != nil {
			return serializer.GetResponse(serializer.ErrorCreateCarinfo)
		}
	}

	return serializer.GetResponse(serializer.Success, map[string]interface{}{"id": parkinfo.ID, "车辆类型": carinfo.CarType})
}
