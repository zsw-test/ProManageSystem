package expressage

import (
	"ProManageSystem/service/expressage"

	"github.com/gin-gonic/gin"
)

func ExpressageCreate(c *gin.Context) {
	service := expressage.ExpressageCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ExpressageCreate()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}

}
