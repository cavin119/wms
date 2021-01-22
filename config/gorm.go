package config

type Mysql struct {
	Path         string `yaml:"path"`
	Config       string `yaml:"config"`
	Dbname       string `yaml:"dbName"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	LogMode      bool   `yaml:"logMode"`
	LogZap       string `yaml:"logZap"`
}