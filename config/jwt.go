package config

type JWT struct {
	SigningKey string `yaml:"signingKey"`
	ExpiresTime int64  `yaml:"expiresTime"`
}
