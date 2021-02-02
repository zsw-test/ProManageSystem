package owner

import (
	"ProManageSystem/service/owner"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OwnerSave(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := owner.OwnerSaveService{
		Id: id,
	}
	if err := c.ShouldBind(&service); err == nil {
		res := service.OwnerSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}

}
