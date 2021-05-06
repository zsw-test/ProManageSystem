package park

import (
	"ProManageSystem/service/carinfo"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CarinfoGet(c *gin.Context) {
	carnumber := c.Param("carnumber")
	service := carinfo.CarinfoGetService{}
	res := service.CarinfoGet(carnumber)
	c.JSON(200, res)
}

func CarinfoGetTotal(c *gin.Context) {
	keyword := c.Query("keyword")
	service := carinfo.CarinfoGetService{}
	res := service.CarinfoGetTotal(keyword)
	c.JSON(200, res)
}

func CarinfoGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	keyword := c.Query("keyword")
	service := carinfo.CarinfoGetService{}
	res := service.CarinfoGetPage(pageindex, pagesize, keyword)
	c.JSON(200, res)
}
