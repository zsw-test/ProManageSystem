package service

import (
	"ProManageSystem/model"
	"fmt"

	"github.com/pkg/errors"
)

type UserService struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Utype    int8   `form:"utype"`
}

func (service *UserService) Userlogin() error {
	user, err := model.GetUser(service.Username)
	if err != nil {
		return err
	}
	if service.Utype == 0 {
		_, err := model.GetOwner(service.Username)
		if err != nil {
			return err
		}
	} else {
		_, err := model.GetManager(service.Username)
		if err != nil {
			return err
		}
	}

	if user == nil {
		return errors.New("没有找到用户")
	}
	if user.Upassword != service.Password {
		fmt.Println("密码错误")
		return errors.New("密码错误")
	}

	return err
}
