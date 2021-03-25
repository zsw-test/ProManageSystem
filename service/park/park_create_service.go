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
			Result: serializer.GetResult(serializer.ErrorParkCreate),
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
		Data:   map[string]interface{}{"id": park.ID},
	}
}
