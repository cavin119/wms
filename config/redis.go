package config

type Redis struct {
	DB       int    `yaml:"db"`
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}