package manager

import (
	"ProManageSystem/service/manager"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ManagerGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := manager.ManagerGetService{
		Id: id,
	}
	res := service.ManagerGet()
	c.JSON(200, res)
}
