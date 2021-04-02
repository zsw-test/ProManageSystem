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

func FaceDetect(c *gin.Context) {
	song := make(map[string]interface{})
	song["image"] = "https://ss2.bdstatic.com/70cFvnSh_Q1YnxGkpoWK1HF6hhy/it/u=3879284015,3445644955&fm=26&gp=0.jpg"
	song["image_type"] = "URL"
	bytesData, err := json.Marshal(song)
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
func Face(c *gin.Context) {

}
