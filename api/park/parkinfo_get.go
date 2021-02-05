package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/parkinfo"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParkInfoGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	service := parkinfo.ParkInfoGetPageService{}
	res := service.ParkInfoGetPage(pageindex, pagesize)
	c.JSON(200, res)
}
func ParkInfoGetTotal(c *gin.Context) {
	service := parkinfo.ParkInfoTotalService{}
	res := service.ParkInfoTotal()
	c.JSON(200, res)
}
