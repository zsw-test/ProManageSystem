package announce

import (
	"ProManageSystem/model/announce"
	"ProManageSystem/serializer"
)

type AnnounceSaveService struct {
	Id int
	// 标题
	Title string
	// 内容
	Text string
	// 发布人
	Managername string
}

func (service *AnnounceSaveService) AnnounceSave() serializer.Response {
	exp, err := announce.GetAnnouncebyid(service.Id)
	if err != nil {
		return serializer.GetResponse(serializer.ErrorGet)
	}
	exp.Title = service.Title
	exp.Text = service.Text
	exp.Managername = service.Managername
	err = exp.Save()
	if err != nil {
		return serializer.GetResponse(serializer.ErrorSave)
	}
	return serializer.GetResponse(serializer.Success)
}
