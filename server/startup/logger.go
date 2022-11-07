package startup

import "github.com/Heqiaomu/goweb/tools/logger"

func InitLogger() {
	z := logger.Logger()
	z.Info("日志组件初始化完成...")
}
