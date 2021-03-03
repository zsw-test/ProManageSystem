package house

import (
	"ProManageSystem/service/house"

	"github.com/gin-gonic/gin"
)

func HouseCreate(c *gin.Context) {
	service := house.HouseCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.HouseCreate()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}

}
