package manager

import (
	"ProManageSystem/model/manager"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
)

type ManagerGetService struct {
	Id int `form:"id"`
}

func (service *ManagerGetService) ManagerGet() serializer.Response {
	manager, err := manager.GetManagerbyid(service.Id)
	if err != nil {
		return serializer.Response{
			Code:   serializer.NotExistUser,
			Result: serializer.GetResult(serializer.NotExistUser),
		}
	} else {
		return serializer.Response{
			Code:   serializer.Success,
			Result: serializer.GetResult(serializer.Success),
			Data:   manager,
		}
	}
}
func (service *ManagerGetService) ManagerGetTotal(keyword string) serializer.Response {
	count, err := manager.GetManagerTotal(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}
