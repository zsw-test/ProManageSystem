package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/parkinfo"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ParkInfoIn(c *gin.Context) {
	carnumber := c.Param("carnumber")
	fmt.Println(carnumber)
	if carnumber == "" {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	service := parkinfo.ParkInfoInService{
		CarNumber: carnumber,
	}
	res := service.ParkInfoIn()
	c.JSON(200, res)
}
