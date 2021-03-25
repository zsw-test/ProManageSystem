package expressage

import (
	"ProManageSystem/model/expressage"
	"ProManageSystem/serializer"
)

type ExpressageSaveService struct {
	Id              int    `form:"id"`
	Ownername       string `form:"ownername"`
	ExpressLocation string `form:"expresslocation"`
	Istake          bool   `form:"istake"`
}

func (service *ExpressageSaveService) ExpressageSave() serializer.Response {
	exp, err := expressage.GetExpressagebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageGet)
	}
	exp.Ownername = service.Ownername
	exp.ExpressLocation = service.ExpressLocation
	exp.Istake = service.Istake
	err = exp.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorExpressageSave)
	}
	return serializer.GetResponse(serializer.Success)
}
