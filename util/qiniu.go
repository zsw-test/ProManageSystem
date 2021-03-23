// package util
// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/qiniu/go-sdk/auth/qbox"
// 	"github.com/qiniu/go-sdk/storage"

// )
// accessKey := "WeLgPYJU2FhvjgOqFp-MfPLgcybp9CSqGJ1hfeAP"
// secretKey := "b_a8AMSaZnqEH5WVnVke1-9WvSWGfq-LVq5SESue"
// bucket:="promanage"
// putPolicy := storage.PutPolicy{
// 		Scope: bucket,
// }
// mac := qbox.NewMac(accessKey, secretKey)

// func GetQiniuToken(c *gin.Context)string{
// 	upToken := putPolicy.UploadToken(mac)
// 	return upToken
// }
