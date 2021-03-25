package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
)

type OwnerGetPageService struct {
}

func (service *OwnerGetPageService) OwnerGetPage(pageindex, pagesize int) serializer.Response {
	ownerlist, err := owner.GetOwnerPage(pageindex, pagesize)
	code := serializer.Success
	if err != nil {
		code = serializer.ErrorGet
		return serializer.Response{
			Code:   code,
			Result: serializer.GetResult(code),
		}
	}
	res := serializer.Response{
		Code:   code,
		Result: serializer.GetResult(code),
		Data:   ownerlist,
	}
	return res

}
