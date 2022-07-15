package cache_manager

import (
	"time"

	"github.com/aditya37/geospatial-tracking/repository"
	goredis "github.com/go-redis/redis/v7"
)

type cachemanager struct {
	client *goredis.Client
}

func CacheManager(
	client *goredis.Client,
) repository.CacheManager {
	return &cachemanager{
		client: client,
	}
}

func (cm *cachemanager) Get(key string) (string, error) {
	return cm.client.Get(key).Result()
}
func (cm *cachemanager) Set(key string, value interface{}, exp time.Duration) error {
	return cm.client.Set(key, value, exp).Err()
}
func (cm *cachemanager) Close() error {
	return cm.client.Close()
}
