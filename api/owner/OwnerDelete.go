package owner

import (
	"ProManageSystem/service/owner"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OwnerDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := owner.OwnerDeleteService{
		Id: id,
	}
	res := service.OwnerDelete()
	c.JSON(200, res)
}
