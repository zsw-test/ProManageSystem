package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/parkinfo"

	"github.com/gin-gonic/gin"
)

func ParkInfoDelete(c *gin.Context) {
	carnumber := c.Param("carnumber")
	if carnumber == "" {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	service := parkinfo.ParkInfoDeleteService{
		CarNumber: carnumber,
	}
	res := service.ParkInfoDelete()
	c.JSON(200, res)
}
