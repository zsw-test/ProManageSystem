package house

import (
	"ProManageSystem/model/house"
	"ProManageSystem/serializer"
)

type ResidentAddService struct {
	HouseId   int
	Address   string
	Residents []house.Resident
}

func (service *ResidentAddService) ResidentAdd() serializer.Response {
	for _, resident := range service.Residents {
		resident.HouseId = service.HouseId
		resident.Address = service.Address
		err := resident.Create()
		if err != nil {
			return serializer.GetResponse(serializer.ErrorAddResident)
		}
	}
	return serializer.GetResponse(serializer.Success)
}
