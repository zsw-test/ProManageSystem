package api

import (
	"ProManageSystem/cache"
	"ProManageSystem/service"
	"fmt"
	"net/http"

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
			session, err := cache.SessionStore.Get(c.Request, "sessionid")
			if err != nil {
				fmt.Println("sessionid 登陆时不存在,即将存入session")
			}
			session.Values["username"] = service.Username
			err = session.Save(c.Request, c.Writer)
			if err != nil {
				fmt.Println(err.Error())
			}
			c.Redirect(http.StatusSeeOther, "/index")
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

func IndexHtml(c *gin.Context) {
	c.HTML(200, "index.html", "")
}
