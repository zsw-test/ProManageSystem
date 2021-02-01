package owner

import (
	"ProManageSystem/service/owner"

	"github.com/gin-gonic/gin"
)

func OwnerLogin(c *gin.Context) {
	service := owner.OwnerLoginService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.OwnerLogin()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
