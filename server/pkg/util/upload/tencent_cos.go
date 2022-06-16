package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/mlsjla/gin-nuxt/server/internal/app/config"
	"github.com/tencentyun/cos-go-sdk-v5"
)

type TencentCOS struct{}

// UploadFile upload file to COS
func (*TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), config.C.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return config.C.TencentCOS.BaseURL + "/" + config.C.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(key string) error {
	client := NewClient()
	name := config.C.TencentCOS.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + config.C.TencentCOS.Bucket + ".cos." + config.C.TencentCOS.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.C.TencentCOS.SecretID,
			SecretKey: config.C.TencentCOS.SecretKey,
		},
	})
	return client
}
