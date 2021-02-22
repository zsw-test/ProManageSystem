package owner

import (
	"ProManageSystem/service/owner"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OwnerGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := owner.OwnerGetService{
		Id: id,
	}
	res := service.OwnerGet()
	c.JSON(200, res)
}
func OwnerGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := owner.OwnerGetPageService{}
	res := service.OwnerGetPage(pageindex, pagesize)
	c.JSON(200, res)
}
func OwnerGetTotal(c *gin.Context) {
	service := owner.OwnerGetService{}
	res := service.OwnerGetTotal()
	c.JSON(200, res)
}
