package charge

import (
	"ProManageSystem/service/charge"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChargeRecordGet(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeRecordGetService{}
	res := service.ChargeRecordGet(houseid)
	c.JSON(200, res)
}

func ChargeRecordGetTotal(c *gin.Context) {
	service := charge.ChargeRecordGetService{}
	res := service.ChargeRecordGetTotal()
	c.JSON(200, res)
}

func ChargeRecordGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeRecordGetService{}
	res := service.ChargeRecordGetPage(pageindex, pagesize)
	c.JSON(200, res)
}
