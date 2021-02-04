package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
)

type OwnerGetService struct {
	Id int `form:"id"`
}

func (service *OwnerGetService) OwnerGet() serializer.Response {
	owner, err := owner.GetOwnerbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.NotExistUser,
			Result: "不存在用户",
		}
	} else {
		return serializer.Response{
			Code:   serializer.Sucess,
			Result: "查询成功",
			Data:   owner,
		}
	}
}
