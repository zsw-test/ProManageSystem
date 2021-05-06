package manager

import (
	"ProManageSystem/service/manager"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ManagerGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("managerid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := manager.ManagerGetService{
		Id: id,
	}
	res := service.ManagerGet()
	c.JSON(200, res)
}

func ManagerGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	keyword := c.Query("keyword")
	service := manager.ManagerGetPageService{}
	res := service.ManagerGetPage(pageindex, pagesize, keyword)
	c.JSON(200, res)
}
func ManagerGetTotal(c *gin.Context) {
	keyword := c.Query("keyword")
	service := manager.ManagerGetService{}
	res := service.ManagerGetTotal(keyword)
	c.JSON(200, res)
}
