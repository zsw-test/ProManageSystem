package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

type OwnerUntieparkService struct {
	Id int `form:"id"`
}

//解除车位绑定
func (service *OwnerUntieparkService) OwnerUntiepark() serializer.Response {
	owner, err := owner.GetOwnerbyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.NotExistUser)
	}
	park, err := parkmodel.GetParkbyid(owner.ParkId)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGetPark)
	}
	park.Ownerid = 0
	err = park.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}

	owner.ParkId = 0
	err = owner.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}
	return serializer.GetResponse(serializer.Sucess)
}
