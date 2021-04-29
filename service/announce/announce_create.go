package announce

import (
	"ProManageSystem/model/announce"
	"ProManageSystem/serializer"
)

type AnnounceCreateService struct {
	// 标题
	Title string
	// 内容
	Text string
	// 发布人
	Managername string
}

func (service *AnnounceCreateService) AnnounceCreate() serializer.Response {
	exp := announce.Announce{
		Title:       service.Title,
		Text:        service.Text,
		Managername: service.Managername,
	}
	err := exp.Create()
	if err != nil {
		return serializer.Response{
			Code:   serializer.ErrorRegister,
			Result: serializer.GetResult(serializer.ErrorRegister),
		}
	}
	return serializer.Response{
		Code:   serializer.Success,
		Result: serializer.GetResult(serializer.Success),
		Data:   map[string]interface{}{"id": exp.ID},
	}
}
