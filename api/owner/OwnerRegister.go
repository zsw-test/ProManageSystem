package owner

import (
	"ProManageSystem/service/owner"

	"github.com/gin-gonic/gin"
)

//业主注册
func OwnerRegister(c *gin.Context) {
	service := owner.OwnerRegisterService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.OwnerRegister()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
