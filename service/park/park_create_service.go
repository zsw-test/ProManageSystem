package park

import (
	"ProManageSystem/model/parkmodel"
	"ProManageSystem/serializer"
)

//车位创建服务
type ParkCreateService struct {
	Ownerid  int    `form:"ownerid"`
	Status   int8   `form:"status"`
	Location string `form:"location"`
}

func (service *ParkCreateService) ParkCreate() serializer.Response {
	park := parkmodel.Park{
		Ownerid:  service.Ownerid,
		Status:   service.Status,
		Location: service.Location,
	}
	err := park.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ExistUser,
			Result: serializer.GetResult(serializer.ExistUser),
		}
	}
	return serializer.Response{
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
		Data:   map[string]interface{}{"id": park.ID},
	}
}
