package auth

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
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
	code := serializer.Sucess
	isExist, err := owner.CheckOwnerAuth(username, password)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if isExist {
		token, err := util.GenerateOwnerToken(username, password)
		if err != nil {
			code = serializer.ErrorCreatToken
		} else {
			data["token"] = token

			code = serializer.Sucess
		}
	} else {
		code = serializer.ErrorAuth
	}

	c.JSON(http.StatusOK, serializer.Response{
		Code:   code,
		Result: "授权生成token",
		Data:   data,
	})
}
