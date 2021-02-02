package jwt

import (
	"ProManageSystem/seralizer"
	"ProManageSystem/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func JWTowner() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = seralizer.Sucess
		token := c.Query("token")
		if token == "" {
			code = seralizer.InvaildParams
		} else {
			claims, err := util.ParseOwnerToken(token)
			if err != nil {
				code = seralizer.ErrorCheckToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = seralizer.ErrorTokenTimeout
			}
		}

		if code != seralizer.Sucess {
			c.JSON(http.StatusUnauthorized, seralizer.Response{
				Code:   code,
				Result: seralizer.GetResult(code),
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

		code = seralizer.Sucess
		token := c.Query("token")
		if token == "" {
			code = seralizer.InvaildParams
		} else {
			claims, err := util.ParseManagerToken(token)
			if err != nil {
				code = seralizer.ErrorCheckToken
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = seralizer.ErrorTokenTimeout
			}
		}

		if code != seralizer.Sucess {
			c.JSON(http.StatusUnauthorized, seralizer.Response{
				Code:   code,
				Result: seralizer.GetResult(code),
				Data:   data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
