package announce

import (
	"ProManageSystem/service/announce"

	"github.com/gin-gonic/gin"
)

func AnnounceCreate(c *gin.Context) {
	service := announce.AnnounceCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.AnnounceCreate()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}

}
