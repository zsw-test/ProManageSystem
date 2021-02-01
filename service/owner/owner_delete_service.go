package owner

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type OwnerDeleteService struct {
	Id int `form:"id"`
}

func (service *OwnerDeleteService) OwnerDelete() seralizer.Response {
	owner, err := model.GetOwnerbyid(service.Id)
	if err != nil {
		return seralizer.Response{
			Code:   "404",
			Result: "用户不存在",
		}
	}
	err = owner.Delete()
	if err != nil {
		return seralizer.Response{
			Code:   "404",
			Result: "删除失败",
		}
	}
	return seralizer.Response{
		Code:   "200",
		Result: "删除成功",
	}
}
