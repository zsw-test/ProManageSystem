package owner

import (
	"ProManageSystem/model/owner"
	"ProManageSystem/serializer"
	"ProManageSystem/util"
)

type OwnerLoginService struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (service *OwnerLoginService) OwnerLogin() serializer.Response {
	data := map[string]interface{}{}
	res := serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
	}
	owner, err := owner.GetOwnerbyname(service.Username)
	if err != nil {
		res.Code = serializer.NotExistUser
		res.Result = serializer.GetResult(res.Code)
		return res
	}
	if service.Password != owner.Password {
		res.Code = serializer.ErrorPassword
		res.Result = serializer.GetResult(res.Code)
		return res
	}

	//登陆成功 生成jwt权限token
	token, err := generateOwnerToken(owner.Username, owner.Password)
	if err != nil {
		res.Code = serializer.ErrorCreatToken
		res.Result = serializer.GetResult(res.Code)
		return res
	}
	data["id"] = owner.ID
	data["token"] = token
	data["houseid"] = owner.Houseid
	data["parkid"] = owner.ParkId
	res.Code = serializer.Success
	res.Result = serializer.GetResult(res.Code)
	res.Data = data
	return res
}

func generateOwnerToken(username, password string) (string, error) {
	var token string
	isExist, err := owner.CheckOwnerAuth(username, password)
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
