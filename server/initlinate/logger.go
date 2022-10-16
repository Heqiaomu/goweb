package initlinate

import "gitee.com/goweb/tools/logger"

func InitLogger() {
	z := logger.Logger()
	z.Info("日志组件初始化完成...")

}
