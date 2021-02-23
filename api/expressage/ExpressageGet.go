package expressage

import (
	"ProManageSystem/service/expressage"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ExpressageGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("expressageid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := expressage.ExpressageGetService{
		Id: id,
	}
	res := service.ExpressageGet()
	c.JSON(200, res)
}

func ExpressageGetTotal(c *gin.Context) {
	service := expressage.ExpressageGetService{}
	res := service.ExpressageGetTotal()
	c.JSON(200, res)
}

func ComplaintGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := expressage.ExpressageGetService{}
	res := service.ExpressageGetPage(pageindex, pagesize)
	c.JSON(200, res)
}
