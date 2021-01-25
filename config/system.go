package config

type System struct {
	Name       string `yaml:"name"`
	Env           string `yaml:"env"`
	Addr          string `yaml:"addr"`
	DbType        string `yaml:"dbType"`
	OssType       string `yaml:"ossType"`
	UseMultipoint bool   `yaml:"useMultipoint"`
}