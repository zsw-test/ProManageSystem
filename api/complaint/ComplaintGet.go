package complaint

import (
	"ProManageSystem/service/complaint"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ComplaintGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := complaint.ComplaintGetService{
		Id: id,
	}
	res := service.ComplaintGet()
	c.JSON(200, res)
}

func ComplaintGetTotal(c *gin.Context) {
	service := complaint.ComplaintGetService{}
	res := service.ComplaintGetTotal()
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
	service := complaint.ComplaintGetService{}
	res := service.ComplaintGetPage(pageindex, pagesize)
	c.JSON(200, res)
}
