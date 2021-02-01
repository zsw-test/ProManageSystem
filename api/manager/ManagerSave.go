package manager

import (
	"ProManageSystem/service/manager"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ManagerSave(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := manager.ManagerSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Id = id
		res := service.ManagerSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
