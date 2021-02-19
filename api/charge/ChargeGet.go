package charge

import (
	"ProManageSystem/service/charge"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChargeGet(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeGetService{}
	res := service.ChargeGet(houseid)
	c.JSON(200, res)
}

func ChargeGetWater(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeGetService{}
	res := service.ChargeGetWater(houseid)
	c.JSON(200, res)
}
func ChargeGetGas(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeGetService{}
	res := service.ChargeGetGas(houseid)
	c.JSON(200, res)
}
func ChargeGetElectric(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeGetService{}
	res := service.ChargeGetElectric(houseid)
	c.JSON(200, res)
}
func ChargeGetProperty(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeGetService{}
	res := service.ChargeGetProperty(houseid)
	c.JSON(200, res)
}

func ChargeGetTotal(c *gin.Context) {
	service := charge.ChargeGetService{}
	res := service.ChargeGetTotal()
	c.JSON(200, res)
}

func ChargeGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeGetService{}
	res := service.ChargeGetPage(pageindex, pagesize)
	c.JSON(200, res)
}
