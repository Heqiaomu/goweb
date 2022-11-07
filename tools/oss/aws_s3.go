package oss

import (
	"errors"
	"fmt"
	"github.com/sun-iot/goweb/config"
	"github.com/sun-iot/goweb/tools/logger"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

type AwsS3 struct{}

func (*AwsS3) UploadFile(file *multipart.FileHeader) (string, string, error) {
	session := newSession()
	uploader := s3manager.NewUploader(session)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := config.GetConfig().OSS.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		logger.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.GetConfig().OSS.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		logger.Error("function uploader.Upload() Filed", zap.Any("err", err.Error()))
		return "", "", err
	}

	return config.GetConfig().OSS.AwsS3.BaseURL + "/" + filename, fileKey, nil
}

//@author: [WqyJh](https://github.com/WqyJh)
//@object: *AwsS3
//@function: DeleteFile
//@description: Delete file from Aws S3 using aws-sdk-go. See https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/s3-example-basic-bucket-operations.html#s3-examples-bucket-ops-delete-bucket-item
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*AwsS3) DeleteFile(key string) error {
	session := newSession()
	svc := s3.New(session)
	filename := config.GetConfig().OSS.AwsS3.PathPrefix + "/" + key
	bucket := config.GetConfig().OSS.AwsS3.Bucket

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		logger.Error("function svc.DeleteObject() Filed", zap.Any("err", err.Error()))
		return errors.New("function svc.DeleteObject() Filed, err:" + err.Error())
	}

	_ = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	return nil
}

// newSession Create S3 session
func newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String(config.GetConfig().OSS.AwsS3.Region),
		Endpoint:         aws.String(config.GetConfig().OSS.AwsS3.Endpoint), //minio在这里设置地址,可以兼容
		S3ForcePathStyle: aws.Bool(config.GetConfig().OSS.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(config.GetConfig().OSS.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			config.GetConfig().OSS.AwsS3.SecretID,
			config.GetConfig().OSS.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
