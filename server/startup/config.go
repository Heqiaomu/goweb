package startup

import (
	"gitee.com/goweb/config"
	"log"
)

// InitConfig 配置解析
func InitConfig() {
	if err := config.NewConfig(); err != nil {
		log.Fatalln(err)
	}
}
