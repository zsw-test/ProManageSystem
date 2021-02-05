package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/parkinfo"

	"github.com/gin-gonic/gin"
)

//Out算出钱
func ParkInfoOut(c *gin.Context) {
	carnumber := c.Param("carnumber")
	if carnumber == "" {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	service := parkinfo.ParkInfoOutService{
		CarNumber: carnumber,
	}
	res := service.ParkInfoOut()
	c.JSON(200, res)
}
