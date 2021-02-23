package expressage

import (
	"ProManageSystem/service/expressage"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExpressageSave(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("expressageid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := expressage.ExpressageSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Id = id
		res := service.ExpressageSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
