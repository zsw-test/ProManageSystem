package complaint

import (
	"ProManageSystem/service/complaint"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ComplaintDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := complaint.ComplaintDeleteService{
		Id: id,
	}
	res := service.ComplaintDelete()
	c.JSON(200, res)
}
