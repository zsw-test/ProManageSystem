package repair

import (
	"ProManageSystem/service/repair"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RepairGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("repairid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := repair.RepairGetService{
		Id: id,
	}
	res := service.RepairGet()
	c.JSON(200, res)
}
func RepairGetOwner(c *gin.Context) {
	service := repair.RepairGetService{}
	service.Ownername = c.Query("Ownername")
	c.JSON(200, service.RepairGetOwner())
}
func RepairGetManager(c *gin.Context) {
	service := repair.RepairGetService{}
	service.Managername = c.Query("Managername")
	c.JSON(200, service.RepairGetManager())
}

func RepairGetTotal(c *gin.Context) {
	service := repair.RepairGetService{}
	res := service.RepairGetTotal()
	c.JSON(200, res)
}

func ComplaintGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := repair.RepairGetService{}
	res := service.RepairGetPage(pageindex, pagesize)
	c.JSON(200, res)
}
