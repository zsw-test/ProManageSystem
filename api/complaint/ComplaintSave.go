package complaint

import (
	"ProManageSystem/service/complaint"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ComplaintSave(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("complaintid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := complaint.ComplaintSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Id = id
		res := service.ComplaintSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
func ComplaintDispatch(c *gin.Context) {
	service := complaint.ComplaintSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ComplaintDispatch()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
func ComplaintResolve(c *gin.Context) {
	service := complaint.ComplaintSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ComplaintResolve()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
