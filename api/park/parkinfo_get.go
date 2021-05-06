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
	keyword := c.Query("keyword")
	service := parkinfo.ParkInfoGetPageService{}
	res := service.ParkInfoGetPage(pageindex, pagesize, keyword)
	c.JSON(200, res)
}
func ParkInfoGetTotal(c *gin.Context) {
	keyword := c.Query("keyword")
	service := parkinfo.ParkInfoTotalService{}
	res := service.ParkInfoTotal(keyword)
	c.JSON(200, res)
}
