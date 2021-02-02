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
			Code:   seralizer.NotExistUser,
			Result: "用户不存在",
		}
	}
	err = owner.Delete()
	if err != nil {
		return seralizer.Response{
			Code:   seralizer.ErrorDelete,
			Result: "删除失败",
		}
	}
	return seralizer.Response{
		Code:   seralizer.Sucess,
		Result: "删除成功",
	}
}
