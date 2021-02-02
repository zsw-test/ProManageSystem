package owner

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type OwnerGetPageService struct {
}

func (service *OwnerGetPageService) OwnerGetPage(pageindex, pagesize int) seralizer.Response {
	ownerlist, err := model.GetOwnerPage(pageindex, pagesize)
	code := seralizer.Sucess
	if err != nil {
		code = seralizer.ErrorGet
		return seralizer.Response{
			Code:   code,
			Result: seralizer.GetResult(code),
		}
	}
	res := seralizer.Response{
		Code:   code,
		Result: seralizer.GetResult(code),
		Data:   ownerlist,
	}
	return res

}
