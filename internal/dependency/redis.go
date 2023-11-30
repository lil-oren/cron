package dependency

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/lil-oren/cron/internal/constant"
)

func NewRedisClient(config Config, logger Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(constant.RedisConnectionTemplate,
			config.RedisCache.HOST,
			config.RedisCache.PORT,
		),
	})

	logger.Infof("Successfully Load Redis Client", nil)

	return client
}
