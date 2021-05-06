package access

import (
	"ProManageSystem/model/accessrecord"
	"ProManageSystem/serializer"
)

type AccessGetService struct {
}

func (service *AccessGetService) AccessGetAll(keyword string) serializer.Response {
	explist, err := accessrecord.GetAccessAll(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, explist)
}
