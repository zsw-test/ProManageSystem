package complaint

import (
	"ProManageSystem/service/complaint"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ComplaintGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("complaintid"))
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
	keyword := c.Query("keyword")
	service := complaint.ComplaintGetService{}
	res := service.ComplaintGetTotal(keyword)
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
	keyword := c.Query("keyword")
	service := complaint.ComplaintGetService{}
	res := service.ComplaintGetPage(pageindex, pagesize, keyword)
	c.JSON(200, res)
}
func ComplaintGetOwner(c *gin.Context) {
	service := complaint.ComplaintGetService{}
	service.Ownername = c.Query("Ownername")
	c.JSON(200, service.ComplaintGetOwner())
}
func ComplaintGetManager(c *gin.Context) {
	service := complaint.ComplaintGetService{}
	service.Managername = c.Query("Managername")
	c.JSON(200, service.ComplaintGetManager())
}
