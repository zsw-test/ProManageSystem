package jwt

import (
	"ProManageSystem/serializer"
	"ProManageSystem/util"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTowner() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = serializer.Sucess
		token := c.GetHeader("Token")
		fmt.Println(c.Cookie("token"))
		if token == "" {
			code = serializer.InvaildParams
		} else {
			claims, err := util.ParseOwnerToken(token)
			if err != nil {
				code = serializer.ErrorCheckToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = serializer.ErrorTokenTimeout
			}
		}

		if code != serializer.Sucess {
			c.JSON(200, serializer.Response{
				Code:   code,
				Result: serializer.GetResult(code),
				Data:   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func JWTmanager() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = serializer.Sucess
		token := c.GetHeader("Token")
		if token == "" {
			code = serializer.InvaildParams
		} else {
			claims, err := util.ParseManagerToken(token)
			if err != nil {
				code = serializer.ErrorCheckToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = serializer.ErrorTokenTimeout
			}
		}

		if code != serializer.Sucess {
			c.JSON(200, serializer.Response{
				Code:   code,
				Result: serializer.GetResult(code),
				Data:   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
