package carinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
	"time"
)

type CarInfoBuyService struct {
	Carnumber string `form:"carnumber"`
	Cartype   string `form:"cartype"`
}

func (service *CarInfoBuyService) CarInfoBuy() serializer.Response {
	//首先查询车辆信息是否存在物业系统中
	carinfo, err := parkmodel.GetCarInfobyCarnumber(service.Carnumber)
	if err != nil {
		//如果不存在直接创建一个购买车信息
		carinfo = &parkmodel.CarInfo{
			CarNumber: service.Carnumber,
			CarType:   service.Cartype,
		}
		err = carinfo.Create()
		if err != nil {
			return serializer.GetResponse(serializer.ErrorCarinfoCreate)
		}
	}
	//如果还没过期  就在过期时间加上去
	var lasttime time.Time
	if carinfo.ExpireTime.After(time.Now()) {
		lasttime = carinfo.ExpireTime
	} else {
		lasttime = time.Now()
	}
	if service.Cartype == parkmodel.YearCar {
		carinfo.CarType = service.Cartype
		carinfo.ExpireTime = lasttime.AddDate(1, 0, 0)
		carinfo.Save()
		return serializer.GetResponse(serializer.Success, carinfo)
	}
	if service.Cartype == parkmodel.MounthCar {
		carinfo.CarType = service.Cartype
		carinfo.ExpireTime = lasttime.AddDate(0, 1, 0)
		carinfo.Save()
		return serializer.GetResponse(serializer.Success, carinfo)
	}
	return serializer.GetResponse(serializer.ErrorCarType)
}
