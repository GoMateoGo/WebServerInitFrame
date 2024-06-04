package config

type Config struct {
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
}
