package announce

import (
	"ProManageSystem/service/announce"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AnnounceDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("announceid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := announce.AnnounceDeleteService{
		Id: id,
	}
	res := service.AnnounceDelete()
	c.JSON(200, res)
}
