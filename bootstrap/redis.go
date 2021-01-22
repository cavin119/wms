package bootstrap

import (
	"7youo-wms/global"
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"os"
)

func Redis() *redis.Client {
	redisConfig := global.G_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
	var ctx = context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		global.G_LOG.Error("redis connect failed", zap.Any("err", err))
		os.Exit(1)
		return nil
	}
	global.G_LOG.Info("redis connect ping response:", zap.Any("pong", pong))
	return client

}
