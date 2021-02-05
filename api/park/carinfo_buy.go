package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/carinfo"

	"github.com/gin-gonic/gin"
)

func CarInfoBuy(c *gin.Context) {
	service := carinfo.CarInfoBuyService{}
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
	}
	res := service.CarInfoBuy()
	c.JSON(200, res)
}
