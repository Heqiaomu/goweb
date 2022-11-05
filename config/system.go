package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                  // 环境值
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                               // 端口值
	DbType        string `mapstructure:"mysql-type" json:"mysql-type" yaml:"mysql-type"`             // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                   // Oss类型
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	TLS           TLS    `mapstructure:"tls" json:"tls" yaml:"tls"`
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
}

type TLS struct {
	Enabled bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	Cert    string `mapstructure:"cert" json:"cert" yaml:"cert"`
	Key     string `mapstructure:"Key" json:"Key" yaml:"Key"`
}
