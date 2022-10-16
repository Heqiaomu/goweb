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

func InitConfig(filePath string) error {
	v = viper.New()
	v.SetConfigFile(filePath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.WatchConfig()
	return nil
}

//v := viper.New()
//v.SetConfigFile(config)
//v.SetConfigType("yaml")
//err := v.ReadInConfig()
//if err != nil {
//panic(any(fmt.Errorf("Fatal error config file: %s \n", err)))
//}
//v.WatchConfig()

func GetViper() *viper.Viper {
	return v
}

func OnConfigChange(f func() error) {
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := f(); err != nil {
			fmt.Printf("Config Watcher Failed.")
		}
	})
}
