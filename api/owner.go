package api

import (
	"ProManageSystem/model"
	"ProManageSystem/service"

	"github.com/gin-gonic/gin"
)

func CreateOwner(c *gin.Context) {
	service := service.OwnerService{}
	if err := c.ShouldBind(&service); err == nil {
		owner := model.Owner{
			Oname:      service.Oname,
			Osex:       service.Osex,
			Oroom:      service.Oroom,
			Otelephone: service.Otelephone,
			User: model.User{
				Username:  service.Username,
				Upassword: service.Upassword,
				Utype:     0,
			},
		}
		err = service.OwnerCreate(&owner)
		if err != nil {
			c.JSON(200, err.Error())
		} else {
			c.JSON(200, "注册成功")
		}
	} else {
		c.JSON(200, err.Error())
	}
}

func SaveOwner(c *gin.Context) {
	service := service.OwnerService{}
	if err := c.ShouldBind(&service); err == nil {
		owner := model.Owner{
			Oname:      service.Oname,
			Osex:       service.Osex,
			Oroom:      service.Oroom,
			Otelephone: service.Otelephone,
			User: model.User{
				Username:  service.Username,
				Upassword: service.Upassword,
				Utype:     0,
			},
		}
		err = service.OwnerSave(&owner)
		if err != nil {
			c.JSON(200, err.Error())
		} else {
			c.JSON(200, "修改成功")
		}
	} else {
		c.JSON(200, err.Error())
	}
}

func GetOwner(c *gin.Context) {
	username := c.Param("username")
	owner, err := model.GetOwner(username)
	if err != nil {
		c.JSON(200, err.Error())
	}
	c.JSON(200, owner)
}
func DeleteOwner(c *gin.Context) {
	username := c.Param("username")
	owner, err := model.GetOwner(username)
	if err != nil {
		c.JSON(200, err.Error())
		return
	}
	err = owner.Delete()
	if err != nil {
		c.JSON(200, err.Error())
	}
	c.JSON(200, "删除成功")
}
