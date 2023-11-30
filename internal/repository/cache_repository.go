package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/lil-oren/cron/internal/constant"
	"github.com/lil-oren/cron/internal/dependency"
	"github.com/lil-oren/cron/internal/dto"
)

type (
	CacheRepository interface {
		SetRecommendedProduct(ctx context.Context, product []dto.HomePageProductResponseBody) error
	}
	cacheRepository struct {
		rd  *redis.Client
		cfg dependency.Config
	}
)

func (r *cacheRepository) SetRecommendedProduct(ctx context.Context, product []dto.HomePageProductResponseBody) error {
	key := constant.RedisRecommendedProductTemplate
	json, err := json.Marshal(product)
	if err != nil {
		return err
	}
	cmd := r.rd.HSetNX(ctx, key, constant.RedisRecommendedProductTemplate, json)
	if err := cmd.Err(); err != nil {
		return err
	}

	return nil
}

func NewCacheRepository(rd *redis.Client, cfg dependency.Config) CacheRepository {
	return &cacheRepository{
		rd:  rd,
		cfg: cfg,
	}
}
