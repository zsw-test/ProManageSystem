package expressage

import (
	"ProManageSystem/model/expressage"
	"ProManageSystem/serializer"
)

type ExpressageCreateService struct {
	Ownername       string `form:"ownername"`
	ExpressLocation string `form:"expresslocation"`
	Telephone       string `form:"telephone"`
	ExpType         string `form:"exptype"`
	Istake          bool   `form:"istake"`
}

func (service *ExpressageCreateService) ExpressageCreate() serializer.Response {
	exp := expressage.Expressage{
		Ownername:       service.Ownername,
		ExpressLocation: service.ExpressLocation,
		Telephone:       service.Telephone,
		ExpType:         service.ExpType,
		Istake:          service.Istake,
	}
	err := exp.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorExpressageCreate,
			Result: serializer.GetResult(serializer.ErrorExpressageCreate),
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
		Data:   map[string]interface{}{"id": exp.ID},
	}
}
