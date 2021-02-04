package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkInfoGetService struct {
	Id int `from:"id"`
}

//get停车信息需要更新停车费
func (service *ParkInfoGetService) ParkInfoGet() serializer.Response {
	parkinfo, err := parkmodel.GetParkInfobyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorNotExistRecord,
			Result: serializer.GetResult(serializer.ErrorNotExistRecord),
		}
	}
	//更新计算费用
	err = parkinfo.UpdateFee()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorUpdateParkFee)
	}
	return serializer.GetResponse(serializer.Sucess, parkinfo)
}
