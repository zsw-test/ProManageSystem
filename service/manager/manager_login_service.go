package manager

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
	"ProManageSystem/util"
)

type ManagerLoginService struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (service *ManagerLoginService) ManagerLogin() seralizer.Response {
	data := map[string]interface{}{}
	res := seralizer.Response{
		Code:   seralizer.Sucess,
		Result: seralizer.GetResult(seralizer.Sucess),
	}
	manager, err := model.GetOwnerbyname(service.Username)
	if err != nil {
		res.Code = seralizer.NotExistUser
		res.Result = seralizer.GetResult(res.Code)
		return res
	}
	if service.Password != manager.Password {
		res.Code = seralizer.ErrorPassword
		res.Result = seralizer.GetResult(res.Code)
		return res
	}

	//登陆成功 生成jwt权限token
	token, err := generateManagerToken(manager.Username, manager.Password)
	if err != nil {
		res.Code = seralizer.ErrorCreatToken
		res.Result = seralizer.GetResult(res.Code)
		return res
	}
	data["id"] = manager.ID
	data["token"] = token
	res.Code = seralizer.Sucess
	res.Result = seralizer.GetResult(res.Code)
	res.Data = data
	return res
}

func generateManagerToken(username, password string) (string, error) {
	var token string
	isExist, err := model.CheckManagerAuth(username, password)
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
