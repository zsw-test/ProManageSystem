package repair

import (
	"ProManageSystem/service/repair"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RepairSave(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("repairid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := repair.RepairSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Id = id
		res := service.RepairSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
func RepairDispatch(c *gin.Context) {
	service := repair.RepairSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.RepairDispatch()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
func RepairResolve(c *gin.Context) {
	service := repair.RepairSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.RepairResolve()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
