package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
	"time"
)

type ParkInfoDeleteService struct {
	CarNumber string `from:"carnumber"`
}

//操作完成 删除车辆和停车信息 从数据库中
//如果是月卡车等  不删除车辆信息
func (service *ParkInfoDeleteService) ParkInfoDelete() serializer.Response {
	parkinfo, err := parkmodel.GetParkInfobyCarnumber(service.CarNumber)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorNotExistRecord)
	}
	parkinfo.Delete()
	carinfo, err := parkmodel.GetCarInfobyCarnumber(service.CarNumber)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorNotExistRecord)
	}
	//如果是临时车  或者到期了就删除车辆存储信息
	if carinfo.CarType == parkmodel.TempoaryCar || time.Now().After(carinfo.ExpireTime) {
		carinfo.Delete()
	}

	return serializer.GetResponse(serializer.Success)
}
