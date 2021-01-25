package global

import (
	"7youo-wms/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	G_VP *viper.Viper
	G_DB *gorm.DB
	G_REDIS *redis.Client
	G_CONFIG *config.Server
	G_LOG *zap.Logger
)
