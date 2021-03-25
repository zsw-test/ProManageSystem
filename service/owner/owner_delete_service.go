package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
)

type OwnerDeleteService struct {
	Id int `form:"id"`
}

func (service *OwnerDeleteService) OwnerDelete() serializer.Response {
	owner, err := owner.GetOwnerbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.NotExistUser,
			Result: "用户不存在",
		}
	}
	err = owner.Delete()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorDelete,
			Result: "删除失败",
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: "删除成功",
	}
}
