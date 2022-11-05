package config

import "gitee.com/goweb/tools/viperfile"

type Server struct {
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email  Email  `mapstructure:"email" json:"email" yaml:"email"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	OSS    OSS    `mapstructure:"oss" json:"oss" yaml:"oss"`
}

var config Server

func NewConfig() error {
	if err := viperfile.GetViper().Unmarshal(&config); err != nil {
		return err
	}
	return nil
}

func GetConfig() *Server {
	return &config
}
