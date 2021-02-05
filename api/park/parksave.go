package park

import (
	"ProManageSystem/serializer"
	"ProManageSystem/service/park"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParkSave(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := park.ParkSaveService{
		Id: id,
	}
	err = c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	res := service.ParkSave()
	c.JSON(200, res)
}
