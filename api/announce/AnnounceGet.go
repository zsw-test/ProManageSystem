package announce

import (
	"ProManageSystem/service/announce"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AnnounceGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("announceid"))
	if err != nil {
		fmt.Println(err.Error())
	}

	service := announce.AnnounceGetService{
		Id: id,
	}
	res := service.AnnounceGet()
	c.JSON(200, res)
}

func AnnounceGetAll(c *gin.Context) {
	service := announce.AnnounceGetService{}
	res := service.AnnounceGetAll()
	c.JSON(200, res)
}
