package manager

import (
	"ProManageSystem/model/manager"
	"ProManageSystem/serializer"
)

type ManagerGetPageService struct {
}

func (service *ManagerGetPageService) ManagerGetPage(pageindex, pagesize int, keyword string) serializer.Response {
	managerlist, err := manager.GetManagerPage(pageindex, pagesize, keyword)
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
		Data:   managerlist,
	}
	return res

}
