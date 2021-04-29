package access

import (
	"ProManageSystem/model/accessrecord"
	"ProManageSystem/serializer"
)

type AccessGetService struct {
}

func (service *AccessGetService) AccessGetAll() serializer.Response {
	explist, err := accessrecord.GetAccessAll()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, explist)
}
