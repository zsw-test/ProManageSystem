package repair

import (
	"ProManageSystem/service/repair"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RepairDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := repair.RepairDeleteService{
		Id: id,
	}
	res := service.RepairDelete()
	c.JSON(200, res)
}
