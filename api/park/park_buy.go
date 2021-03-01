package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/park"

	"github.com/gin-gonic/gin"
)

func ParkBuy(c *gin.Context) {

	service := park.ParkBuyService{}
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	res := service.ParkBuy()
	c.JSON(200, res)
}
