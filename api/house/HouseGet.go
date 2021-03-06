package house

import (
	"ProManageSystem/service/house"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HouseGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.HouseGetService{
		Id: id,
	}
	res := service.HouseGet()
	c.JSON(200, res)
}

func HouseGetTotal(c *gin.Context) {
	keyword := c.Query("keyword")
	service := house.HouseGetService{}
	res := service.HouseGetTotal(keyword)
	c.JSON(200, res)
}

func HouseGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	keyword := c.Query("keyword")
	service := house.HouseGetService{}
	res := service.HouseGetPage(pageindex, pagesize, keyword)
	c.JSON(200, res)
}
