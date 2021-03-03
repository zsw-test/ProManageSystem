package house

import (
	"ProManageSystem/service/house"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HouseSave(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.HouseSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Id = id
		res := service.HouseSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
