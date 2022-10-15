// Package viperfile
/**
 * @Author: sunyang
 * @Email: sunyang@hyperchain.cn
 * @Date: 2022/10/12
 */
package viperfile

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var v *viper.Viper

func init() {
	v = viper.New()
}
func InitConfig(filePath string) error {
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.WatchConfig()
	return nil
}

func GetViperfile() *viper.Viper {
	return v
}

func OnConfigChange(f func() error) {
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := f(); err != nil {
			fmt.Printf("Config Watcher Failed.")
		}
	})
}
