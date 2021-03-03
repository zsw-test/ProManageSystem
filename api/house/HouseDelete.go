package house

import (
	"ProManageSystem/service/house"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HouseDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.HouseDeleteService{
		Id: id,
	}
	res := service.HouseDelete()
	c.JSON(200, res)
}
