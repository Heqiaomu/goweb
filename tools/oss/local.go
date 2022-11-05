package oss

import (
	"errors"
	"gitee.com/goweb/config"
	md5File "gitee.com/goweb/tools/file"
	"gitee.com/goweb/tools/logger"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Local struct{}

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = md5File.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(config.GetConfig().OSS.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		logger.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := config.GetConfig().OSS.Local.StorePath + "/" + filename
	filepath := config.GetConfig().OSS.Local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		logger.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		logger.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		logger.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return filepath, filename, nil
}

func (*Local) DeleteFile(key string) error {
	p := config.GetConfig().OSS.Local.StorePath + "/" + key
	if strings.Contains(p, config.GetConfig().OSS.Local.StorePath) {
		//if err := os.Remove(p); err != nil {
		//	return errors.New("本地文件删除失败, err:" + err.Error())
		//}
	}
	return nil
}
