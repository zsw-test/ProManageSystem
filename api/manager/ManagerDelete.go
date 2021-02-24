package manager

import (
	"ProManageSystem/service/manager"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ManagerDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("managerid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := manager.ManagerDeleteService{
		Id: id,
	}
	res := service.ManagerDelete()
	c.JSON(200, res)
}
