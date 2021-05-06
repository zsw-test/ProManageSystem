package access

import (
	"ProManageSystem/service/access"

	"github.com/gin-gonic/gin"
)

func AccessCreate(c *gin.Context) {
	service := access.AccessCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.AccessCreate()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
func AccessGetAll(c *gin.Context) {
	keyword := c.Query("keyword")
	service := access.AccessGetService{}
	res := service.AccessGetAll(keyword)
	c.JSON(200, res)
}
