package owner

import (
	"ProManageSystem/model"
	"ProManageSystem/seralizer"
	"ProManageSystem/util"
)

type OwnerLoginService struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (service *OwnerLoginService) OwnerLogin() seralizer.Response {
	data := map[string]interface{}{}
	res := seralizer.Response{
		Code:   seralizer.Sucess,
		Result: seralizer.GetResult(seralizer.Sucess),
	}
	owner, err := model.GetOwnerbyname(service.Username)
	if err != nil {
		res.Code = seralizer.NotExistUser
		res.Result = seralizer.GetResult(res.Code)
		return res
	}
	if service.Password != owner.Password {
		res.Code = seralizer.ErrorPassword
		res.Result = seralizer.GetResult(res.Code)
		return res
	}

	//登陆成功 生成jwt权限token
	token, err := generateOwnerToken(owner.Username, owner.Password)
	if err != nil {
		res.Code = seralizer.ErrorCreatToken
		res.Result = seralizer.GetResult(res.Code)
		return res
	}
	data["id"] = owner.ID
	data["token"] = token
	res.Code = seralizer.Sucess
	res.Result = seralizer.GetResult(res.Code)
	res.Data = data
	return res
}

func generateOwnerToken(username, password string) (string, error) {
	var token string
	isExist, err := model.CheckOwnerAuth(username, password)
	if err != nil {
		return token, err
	}
	if isExist {
		token, err = util.GenerateOwnerToken(username, password)
		if err != nil {
			return token, err
		}
	}
	return token, err

}
