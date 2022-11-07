package startup

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sun-iot/goweb/tools/viperfile"
	"log"
)

const (
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)

func InitViper() {
	var config string
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	if config == "" {
		switch gin.Mode() {
		case gin.DebugMode:
			config = ConfigDebugFile
		case gin.ReleaseMode:
			config = ConfigReleaseFile
		case gin.TestMode:
			config = ConfigTestFile
		default:
			config = ConfigDefaultFile
		}
	}
	// 开始进行文件读取
	if err := viperfile.InitConfig(config); err != nil {
		log.Fatalln(err)
	}
}
