package parkinfo

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type ParkInfoSaveService struct {
	CarNumber string `form:"carnumber"`
	Fee       int    `form:"fee"`
}

func (service *ParkInfoSaveService) ParkInfoSave() serializer.Response {
	parkinfo := parkmodel.ParkInfo{
		CarNumber: service.CarNumber,
		Fee:       service.Fee,
	}
	err := parkinfo.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}
	return serializer.GetResponse(serializer.Sucess)
}
