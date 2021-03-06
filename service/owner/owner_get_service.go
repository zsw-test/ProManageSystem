package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"

	"github.com/gin-gonic/gin"
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
			Code:   serializer.Success,
			Result: "查询成功",
			Data:   owner,
		}
	}
}
func (service *OwnerGetService) OwnerGetTotal(keyword string) serializer.Response {
	count, err := owner.GetOwnerTotal(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, gin.H{"count": count})
}
