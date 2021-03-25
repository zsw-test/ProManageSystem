package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"
)

type ReisdentSaveService struct {
	Id      int
	Name    string
	Age     int
	Sex     string
	Work    string
	IdCard  string
	HouseId int
	Address string
}

func (service *ReisdentSaveService) ReisdentSave() serializer.Response {
	resident, err := house.GetResidentbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetResident)
	}
	resident.Name = service.Name
	resident.Age = service.Age
	resident.Sex = service.Sex
	resident.Work = service.Work
	resident.IdCard = service.IdCard
	resident.HouseId = service.HouseId
	resident.Address = service.Address
	err = resident.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSaveResident)
	}
	return serializer.GetResponse(serializer.Success)
}
