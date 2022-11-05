package config

type HuaWeiObs struct {
	Enabled   bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Path      string `mapstructure:"path" json:"path" yaml:"path"`
	Bucket    string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Endpoint  string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	AccessKey string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`
	SecretKey string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`
}
