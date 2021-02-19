package charge

import (
	"ProManageSystem/service/charge"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChargeSave(c *gin.Context) {
	houseid, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := charge.ChargeSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Houseid = houseid
		res := service.ChargeSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
