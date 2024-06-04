package config

type Server struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 端口
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"` // 模式
}
