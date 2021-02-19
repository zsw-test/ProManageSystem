package charge

import (
	"ProManageSystem/service/charge"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChargeWaterPay(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargePayWaterService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Houseid = houseid
		res := service.ChargePayWater()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
func ChargeElectricPay(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargePayElectricService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Houseid = houseid
		res := service.ChargePayElectric()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
func ChargeGasPay(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargePayGasService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Houseid = houseid
		res := service.ChargePayGas()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
func ChargePropertyPay(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargePayPropertyService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Houseid = houseid
		res := service.ChargePayProperty()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
