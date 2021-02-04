package jwt

import (
	"ProManageSystem/serializer"
	"ProManageSystem/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTowner() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = serializer.Sucess
		token := c.Query("token")
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
			c.JSON(http.StatusUnauthorized, serializer.Response{
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
		token := c.Query("token")
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
			c.JSON(http.StatusUnauthorized, serializer.Response{
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
