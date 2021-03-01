package park

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
	"time"
)

type ParkBuyService struct {
	ParkId    int    `form:"parkid"`
	Ownerid   int    `form:"ownerid"`
	Carnumber string `form:"carnumber"`
}

//购买车位业务
func (service *ParkBuyService) ParkBuy() serializer.Response {
	owner, err := owner.GetOwnerbyid(service.Ownerid)
	if err != nil {
		return serializer.GetResponse(serializer.NotExistUser)
	}
	if owner.ParkId != 0 {
		return serializer.GetResponse(serializer.ErrorParkBuy)
	}
	park, err := parkmodel.GetParkbyid(service.ParkId)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	park.Ownerid = service.Ownerid
	err = park.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}
	owner.ParkId = int(park.ID)
	err = owner.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}
	carinfo := &parkmodel.CarInfo{
		CarNumber:  service.Carnumber,
		CarType:    parkmodel.PersonalCar,
		ExpireTime: time.Date(2100, 02, 27, 17, 30, 20, 20, time.Local),
	}
	err = carinfo.Create()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorCarinfoCreate)
	}

	return serializer.GetResponse(serializer.Sucess)
}
