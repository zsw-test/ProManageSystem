package manager

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
)

type ManagerGetPageService struct {
}

func (service *ManagerGetPageService) ManagerGetPage(pageindex, pagesize int) seralizer.Response {
	managerlist, err := model.GetManagerPage(pageindex, pagesize)
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
		Data:   managerlist,
	}
	return res

}
