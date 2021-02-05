package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/park"

	"github.com/gin-gonic/gin"
)

func ParkCreate(c *gin.Context) {
	service := park.ParkCreateService{}
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	res := service.ParkCreate()
	c.JSON(200, res)
}
