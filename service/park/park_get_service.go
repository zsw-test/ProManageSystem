package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ParkGetService struct {
	Id int `from:"id"`
}

func (service *ParkGetService) ParkGet() serializer.Response {
	park, err := parkmodel.GetParkbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorGetPark,
			Result: serializer.GetResult(serializer.ErrorGetPark),
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
		Data:   park,
	}
}

func (service *ParkGetService) ParkGetTotal(keyword string) serializer.Response {
	count, err := parkmodel.GetParkTotal(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}

func (service *ParkGetService) ParkGetFreeList() serializer.Response {
	parklist, err := parkmodel.GetParkFreeList()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	return serializer.GetResponse(serializer.Success, parklist)
}

//获取空闲车位数量   所有车位数量减去在场车辆信息数量  停车信息相当于一个停车记录 只有出去才会删除
func (service *ParkGetService) ParkGetAllCount() serializer.Response {
	parkinfocount, err := parkmodel.GetParkInfoTotal("")
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	parkcount, err := parkmodel.GetParkTotal("")
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	mp := map[string]int{}
	mp["parkinfocount"] = parkinfocount
	mp["parkcount"] = parkcount
	return serializer.GetResponse(serializer.Success, mp)
}
