package oss

import (
	"context"
	"errors"
	"fmt"
	"github.com/sun-iot/goweb/config"
	"github.com/sun-iot/goweb/tools/logger"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

type TencentCOS struct{}

// UploadFile oss file to COS
func (*TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		logger.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), config.GetConfig().OSS.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(any(err))
	}
	return config.GetConfig().OSS.TencentCOS.BaseURL + "/" + config.GetConfig().OSS.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(key string) error {
	client := NewClient()
	name := config.GetConfig().OSS.TencentCOS.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		logger.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + config.GetConfig().OSS.TencentCOS.Bucket + ".cos." + config.GetConfig().OSS.TencentCOS.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.GetConfig().OSS.TencentCOS.SecretID,
			SecretKey: config.GetConfig().OSS.TencentCOS.SecretKey,
		},
	})
	return client
}
