package config

type Server struct {
	Mysql  Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	System System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
