package api

import (
	"ProManageSystem/model"
	"ProManageSystem/service"

	"github.com/gin-gonic/gin"
)

func SaveManager(c *gin.Context) {
	service := service.ManagerService{}
	if err := c.ShouldBind(&service); err == nil {
		manager := model.Manager{
			Mname: service.Mname,
			Msex:  service.Msex,
			Mno:   service.Mno,
			User: model.User{
				Username:  service.Username,
				Upassword: service.Password,
				Utype:     1,
			},
		}
		err = service.ManagerSave(&manager)
		if err != nil {
			c.JSON(200, err.Error())
		} else {
			c.JSON(200, "修改成功")
		}
	} else {
		c.JSON(200, err.Error())
	}
}
func CreateManager(c *gin.Context) {
	service := service.ManagerService{}
	if err := c.ShouldBind(&service); err == nil {
		manager := model.Manager{
			Mname: service.Mname,
			Msex:  service.Msex,
			Mno:   service.Mno,
			User: model.User{
				Username:  service.Username,
				Upassword: service.Password,
				Utype:     1,
			},
		}
		err = service.ManagerCreate(&manager)
		if err != nil {
			c.JSON(200, err.Error())
		} else {
			c.JSON(200, "注册成功")
		}
	} else {
		c.JSON(200, err.Error())
	}
}

func GetManager(c *gin.Context) {
	username := c.Param("username")
	manager, err := model.GetManager(username)
	if err != nil {
		c.JSON(200, err.Error())
	}
	c.JSON(200, manager)
}
func DeleteManager(c *gin.Context) {
	username := c.Param("username")
	manager, err := model.GetManager(username)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}
	err = manager.Delete()
	if err != nil {
		c.JSON(200, err.Error())
	}
	c.JSON(200, "删除成功")
}
