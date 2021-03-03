package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"
)

type HouseCreateService struct {
	//楼栋号码
	Building int
	//单元号码
	Unit int
	//门牌号吗
	Door         int
	Area         float64
	Prorityright int
	HouseType    string
	Ownerid      int
}

func (service *HouseCreateService) HouseCreate() serializer.Response {
	house := house.House{
		Building:     service.Building,
		Unit:         service.Unit,
		Door:         service.Door,
		Area:         service.Area,
		Prorityright: service.Prorityright,
		HouseType:    service.HouseType,
		Ownerid:      service.Ownerid,
	}
	err := house.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorHouseCreate,
			Result: serializer.GetResult(serializer.ErrorHouseCreate),
		}
	}
	return serializer.Response{
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
		Data:   map[string]interface{}{"id": house.ID},
	}
}
