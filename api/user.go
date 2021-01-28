package api

import (
	"ProManageSystem/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	service := service.UserService{}
	if err := c.ShouldBind(&service); err == nil {
		fmt.Println(service)
		err = service.Userlogin()
		//逻辑
		if err != nil {
			c.JSON(200, err.Error())
		} else {
			c.HTML(200, "index.html", gin.H{
				"username": service.Username,
			})
			c.JSON(200, "登陆成功 欢迎您"+service.Username)
		}
	} else {
		c.JSON(200, err.Error())
	}
}

func LoginHtml(c *gin.Context) {
	c.HTML(200, "login.html", "")
}

func RegisterHtml(c *gin.Context) {
	c.HTML(200, "register.html", "")
}
