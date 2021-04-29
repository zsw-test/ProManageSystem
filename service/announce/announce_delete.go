package announce

import (
	"ProManageSystem/model/announce"
	"ProManageSystem/serializer"
)

type AnnounceDeleteService struct {
	Id int `form:"id"`
}

func (service *AnnounceDeleteService) AnnounceDelete() serializer.Response {
	complaint, err := announce.GetAnnouncebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	err = complaint.Delete()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorDelete)
	}
	return serializer.GetResponse(serializer.Success)
}
