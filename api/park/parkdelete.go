package park

import (
	"ProManageSystem/service/park"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParkDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := park.ParkDeleteService{
		Id: id,
	}
	res := service.ParkDelete()
	c.JSON(200, res)
}
