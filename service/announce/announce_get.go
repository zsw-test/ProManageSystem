package announce

import (
	"ProManageSystem/model/announce"
	"ProManageSystem/serializer"
)

type AnnounceGetService struct {
	Id int `form:"id"`
}

func (service *AnnounceGetService) AnnounceGetAll(keyword string) serializer.Response {
	explist, err := announce.GetAnnounceAll(keyword)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, explist)
}

func (service *AnnounceGetService) AnnounceGet() serializer.Response {
	exp, err := announce.GetAnnouncebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	return serializer.GetResponse(serializer.Success, exp)
}
