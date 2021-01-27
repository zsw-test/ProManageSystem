package api

import (
	"ProManageSystem/model"
	"ProManageSystem/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	service := service.UserService{}

	if err := c.ShouldBind(&service); err == nil {
		fmt.Println(service)
		user := &model.User{
			Username: service.Username,
			Password: service.Password,
			Address:  service.Address,
		}
		service.Create(user)
		c.JSON(200, user)
	} else {
		c.JSON(200, err)
	}

}
func GetUser(c *gin.Context) {

}
func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
