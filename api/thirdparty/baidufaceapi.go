package thirdparty

import (
	"ProManageSystem/serializer"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	appid     = 15809152
	apikey    = "dFxqUzN7nC2yZiZMxFB9wrbu"
	secretkey = "xIWMmr5yCLGMRtUkbuPBHEe4bhA9uukI"
)

func AccessToken() string {
	params := url.Values{}

	Url, err := url.Parse("https://aip.baidubce.com/oauth/2.0/token")
	if err != nil {
		panic(err.Error())

	}
	params.Set("grant_type", "client_credentials")
	params.Set("client_id", apikey)
	params.Set("client_secret", secretkey)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	type TokenJson struct {
		RefreshToken  string `json:"refresh_token"`
		ExpiresIn     int64  `json:"expires_in"`
		Scope         string `json:"scope"`
		SessionKey    string `json:"session_key"`
		AccessToken   string `json:"access_token"`
		SessionSecret string `json:"session_secret"`
	}
	var t = &TokenJson{}
	json.Unmarshal(s, t)
	return t.AccessToken
}

type FaceAddService struct {
	Image    string `json:"image"`
	Username string `json:"username"`
}

func FaceDetect(c *gin.Context) {
	params := make(map[string]interface{})
	image := c.PostForm("image")
	params["image"] = image + "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=3879284015,3445644955&fm=26&gp=0.jpg"
	params["image_type"] = "URL"
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "https://aip.baidubce.com/rest/2.0/face/v3/detect?access_token=" + AccessToken()
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := string(respBytes)
	c.JSON(200, serializer.GetResponse(serializer.Success, str))

}
func FaceAdd(c *gin.Context) {
	params := make(map[string]interface{})
	service := FaceAddService{}
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	params["image"] = service.Image
	params["image_type"] = "URL"
	params["group_id"] = "promanage"
	params["user_id"] = service.Username
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add?access_token=" + AccessToken()
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := string(respBytes)
	c.JSON(200, serializer.GetResponse(serializer.Success, str))
}

type FaceSearchService struct {
	Image    string `json:"image"`
	Username string `json:"username"`
}

func FaceSearch(c *gin.Context) {
	params := make(map[string]interface{})
	service := FaceSearchService{}
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(200, serializer.GetResponse(serializer.InvaildParams))
		return
	}
	params["image"] = service.Image
	params["image_type"] = "URL"
	params["group_id_list"] = "promanage"
	params["user_id"] = service.Username
	bytesData, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(bytesData)
	url := "https://aip.baidubce.com/rest/2.0/face/v3/search?access_token=" + AccessToken()
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	str := string(respBytes)
	c.JSON(200, serializer.GetResponse(serializer.Success, str))
}
