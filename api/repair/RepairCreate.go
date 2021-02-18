package repair

import (
	"ProManageSystem/service/repair"

	"github.com/gin-gonic/gin"
)

func RepairCreate(c *gin.Context) {
	service := repair.RepairCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.RepairCreate()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}

}
