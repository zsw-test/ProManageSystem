package thirdparty

import (
	"ProManageSystem/serializer"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

const (
	accessKey = "WeLgPYJU2FhvjgOqFp-MfPLgcybp9CSqGJ1hfeAP"
	secretKey = "b_a8AMSaZnqEH5WVnVke1-9WvSWGfq-LVq5SESue"
	bucket    = "promanage1"
)

func GetQiniuToken(c *gin.Context) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	c.JSON(200, serializer.GetResponse(serializer.Success, upToken))
}

// 上传图片到七牛云接口 客户端直接拼接域名+图片名字即可获取图片链接
func UpPhoto(c *gin.Context) {
	fileheader, _ := c.FormFile("file")
	file, _ := fileheader.Open()
	defer file.Close()
	fmt.Println(fileheader.Filename)

	var data []byte
	buf := make([]byte, 1024)

	for {
		//从file读取到buf中
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read buf fail", err)
			return
		}
		//说明读取结束
		if n == 0 {
			break
		}
		//读取到最终的缓冲区中
		data = append(data, buf[:n]...)
	}

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, fileheader.Filename, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		fmt.Println("puterror:", err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
	// Upload the file to specific dst.
	c.JSON(http.StatusOK, serializer.GetResponse(serializer.Success, fileheader.Filename))

}
