package manager

import (
	"ProManageSystem/service/manager"

	"github.com/gin-gonic/gin"
)

func ManagerLogin(c *gin.Context) {
	service := manager.ManagerLoginService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ManagerLogin()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
