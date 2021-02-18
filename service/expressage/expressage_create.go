package expressage

import (
	"ProManageSystem/model/expressage"
	"ProManageSystem/serializer"
)

type ExpressageCreateService struct {
	Ownername       string `form:"ownername"`
	ExpressLocation string `form:"expresslocation"`
	Istake          bool   `form:"istake"`
}

func (service *ExpressageCreateService) ExpressageCreate() serializer.Response {
	exp := expressage.Expressage{
		Ownername:       service.Ownername,
		ExpressLocation: service.ExpressLocation,
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
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
		Data:   map[string]interface{}{"id": exp.ID},
	}
}
