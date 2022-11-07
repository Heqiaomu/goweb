package startup

import (
	"github.com/sun-iot/goweb/config"
	"log"
)

// InitConfig 配置解析
func InitConfig() {
	if err := config.NewConfig(); err != nil {
		log.Fatalln(err)
	}
}
