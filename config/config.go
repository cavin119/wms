package config

type Server struct {
	System System `yaml:"system"`
	JWT    JWT    `yaml:"jwt"`
	//gorm
	Mysql  Mysql  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	//oss
	Local Local `yaml:"local"`
	Zap Zap `yaml:"zap"`
}