package auth

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
	"ProManageSystem/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string
	Password string
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	data := make(map[string]interface{})
	code := seralizer.Sucess
	isExist, err := model.CheckOwnerAuth(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if isExist {
		token, err := util.GenerateOwnerToken(username, password)
		if err != nil {
			code = seralizer.ErrorCreatToken
		} else {
			data["token"] = token

			code = seralizer.Sucess
		}
	} else {
		code = seralizer.ErrorAuth
	}

	c.JSON(http.StatusOK, seralizer.Response{
		Code:   code,
		Result: "授权生成token",
		Data:   data,
	})
}
