package manager

import (
	"ProManageSystem/service/manager"

	"github.com/gin-gonic/gin"
)

func ManagerRegister(c *gin.Context) {
	service := manager.ManagerRegisterService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ManagerRegister()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
	
}
