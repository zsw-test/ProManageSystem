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
	service := carinfo.CarinfoGetService{}
	res := service.CarinfoGetTotal()
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
	service := carinfo.CarinfoGetService{}
	res := service.CarinfoGetPage(pageindex, pagesize)
	c.JSON(200, res)
}
