package expressage

import (
	"ProManageSystem/service/expressage"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExpressageDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := expressage.ExpressageDeleteService{
		Id: id,
	}
	res := service.ExpressageDelete()
	c.JSON(200, res)
}
