package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkInfoCreateService struct {
	//占用的车位 将这个车位状态改变
	ParkId int `form:"parkid"`
	//车辆号码
	CarNumber string `form:"carnumber"`
	//费用  根据当前时间和创建时间来算出总费用
	Fee int `form:"fee"`
}

//创建一个停车信息同时将车位状态改变
func (service *ParkInfoCreateService) ParkInfoCreate() serializer.Response {
	park, err := parkmodel.GetParkbyid(service.ParkId)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}

	parkinfo := parkmodel.ParkInfo{
		ParkId:    service.ParkId,
		CarNumber: service.CarNumber,
		Fee:       service.Fee,
	}
	err = parkinfo.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorCreateParkinfo)
	}
	park.Status = parkmodel.Occupy
	return serializer.GetResponse(serializer.Sucess, map[string]interface{}{"id": parkinfo.ID})
}
