package manager

import (
	"ProManageSystem/model/manager"
	"ProManageSystem/serializer"
	"ProManageSystem/util"
)

type ManagerLoginService struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (service *ManagerLoginService) ManagerLogin() serializer.Response {
	data := map[string]interface{}{}
	res := serializer.Response{
		Code:   serializer.Sucess,
		Result: serializer.GetResult(serializer.Sucess),
	}
	manager, err := manager.GetManagerbyname(service.Username)
	if err != nil {
		res.Code = serializer.NotExistUser
		res.Result = serializer.GetResult(res.Code)
		return res
	}
	if service.Password != manager.Password {
		res.Code = serializer.ErrorPassword
		res.Result = serializer.GetResult(res.Code)
		return res
	}

	//登陆成功 生成jwt权限token
	token, err := generateManagerToken(manager.Username, manager.Password)
	if err != nil {
		res.Code = serializer.ErrorCreatToken
		res.Result = serializer.GetResult(res.Code)
		return res
	}
	data["id"] = manager.ID
	data["token"] = token
	res.Code = serializer.Sucess
	res.Result = serializer.GetResult(res.Code)
	res.Data = data
	return res
}

func generateManagerToken(username, password string) (string, error) {
	var token string
	isExist, err := manager.CheckManagerAuth(username, password)
	if err != nil {
		return token, err
	}
	if isExist {
		token, err = util.GenerateManagerToken(username, password)
		if err != nil {
			return token, err
		}
	}
	return token, err

}
