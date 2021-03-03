package house

import (
	"ProManageSystem/service/house"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//获取一个住户的信息
func ResidentGet(c *gin.Context) {
	residnetid, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.ResidnetGetService{
		ResidentId: residnetid,
	}
	res := service.ResidentGetByid()
	c.JSON(200, res)
}

//获取小区所有的住户数量
func ResidentGetTotal(c *gin.Context) {
	service := house.ResidnetGetService{}
	res := service.ResidentGetTotal()
	c.JSON(200, res)
}

//获取小区住户分页数据
func ResidentGetPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.ResidnetGetService{}
	res := service.ResidentGetPage(pageindex, pagesize)
	c.JSON(200, res)
}

//获取该房屋的住户数量
func ResidentGetByhouseidTotal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.ResidnetGetService{
		Houseid: id,
	}
	res := service.ResidentGetByhouseidTotal()
	c.JSON(200, res)
}

//获取房屋住户分页数据
func ResidentGetByhouseidPage(c *gin.Context) {
	pageindex, err := strconv.Atoi(c.Query("pageindex"))
	if err != nil {
		fmt.Println(err.Error())
	}
	pagesize, err := strconv.Atoi(c.Query("pagesize"))
	if err != nil {
		fmt.Println(err.Error())
	}
	id, err := strconv.Atoi(c.Param("houseid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.ResidnetGetService{
		Houseid: id,
	}
	res := service.ResidentGetByhouseidPage(pageindex, pagesize)
	c.JSON(200, res)
}

//添加房屋住户
func ResidentAdd(c *gin.Context) {
	service := house.ResidentAddService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ResidentAdd()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

//删除房屋住户
func ResidentDel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("residentid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.ResidentDeleteService{}
	if err := c.ShouldBind(&service); err == nil {
		service.ResidentId = id
		res := service.ResidentDelete()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}

func ResidentSave(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("residentid"))
	if err != nil {
		fmt.Println(err.Error())
	}
	service := house.ReisdentSaveService{}
	if err := c.ShouldBind(&service); err == nil {
		service.Id = id
		res := service.ReisdentSave()
		c.JSON(200, res)
	} else {
		c.JSON(200, err.Error())
	}
}
