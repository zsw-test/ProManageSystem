package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/park"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParkGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("pakrid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := park.ParkGetService{
		Id: id,
	}
	res := service.ParkGet()
	c.JSON(200, res)
}
func ParkGetPage(c *gin.Context) {
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
	service := park.ParkGetPageService{}
	res := service.ParkGetPage(pageindex, pagesize, keyword)
	c.JSON(200, res)
}
func ParkGetTotal(c *gin.Context) {
	keyword := c.Query("keyword")
	service := park.ParkGetService{}
	res := service.ParkGetTotal(keyword)
	c.JSON(200, res)
}
func ParkGetFreeList(c *gin.Context) {
	service := park.ParkGetService{}
	res := service.ParkGetFreeList()
	c.JSON(200, res)
}
func ParkGetAllCount(c *gin.Context) {
	service := park.ParkGetService{}
	res := service.ParkGetAllCount()
	c.JSON(200, res)
}
