package complaint

import (
	"ProManageSystem/service/complaint"

	"github.com/gin-gonic/gin"
)

func ComplaintCreate(c *gin.Context) {
	service := complaint.ComplaintCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ComplaintCreate()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}

}
