package service

import (
	"ProManageSystem/model"
	"fmt"
)

type UserService struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Address  string `form:"address"`
}

func (service *UserService) Get() *model.User {
	user, err := model.GetUser(service.Username)
	if err != nil {
		fmt.Println(err)
	}
	return user
}
func (service *UserService) Create(user *model.User) {
	err := user.Create()
	if err != nil {
		fmt.Println(err)
	}
}
func (service *UserService) Save(user *model.User) {
	err := user.Save()
	if err != nil {
		fmt.Println(err)
	}
}
func (service *UserService) Delete(user *model.User) {
	err := user.Delete()
	if err != nil {
		fmt.Println(err)
	}
}
