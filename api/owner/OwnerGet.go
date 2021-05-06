package owner

import (
	"ProManageSystem/service/owner"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func OwnerGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("ownerid"))
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
	keyword := c.Query("keyword")
	service := owner.OwnerGetPageService{}
	res := service.OwnerGetPage(pageindex, pagesize, keyword)
	c.JSON(200, res)
}
func OwnerGetTotal(c *gin.Context) {
	keyword := c.Query("keyword")
	service := owner.OwnerGetService{}
	res := service.OwnerGetTotal(keyword)
	c.JSON(200, res)
}
