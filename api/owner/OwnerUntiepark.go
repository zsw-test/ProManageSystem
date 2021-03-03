package owner

import (
	"ProManageSystem/service/owner"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OwnerUntiepark(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ownerid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := owner.OwnerUntieparkService{
		Id: id,
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.OwnerUntiepark()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}

}
