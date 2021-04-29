package announce

import (
	"ProManageSystem/service/announce"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AnnounceSave(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("announceid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := announce.AnnounceSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Id = id
		res := service.AnnounceSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
