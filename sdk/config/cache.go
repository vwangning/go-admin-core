package config

import (
	"context"
	"github.com/go-admin-team/go-admin-core/storage"
	"github.com/go-admin-team/go-admin-core/storage/cache"
	"github.com/go-redis/redis/v9"
)

type Cache struct {
	Redis  *RedisConnectOptions
	Memory interface{}
}

// CacheConfig cache配置
var CacheConfig = new(Cache)

func getClusterRedisClient(redisOptions *RedisConnectOptions) (*redis.ClusterClient, error) {

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    redisOptions.Addrs,
		Password: redisOptions.Password,
	})
	ctx := context.Background()
	err := clusterClient.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return clusterClient, nil
}

// Setup 构造cache 顺序 集群redis >redis > 其他 > memory
func (e Cache) Setup() (storage.AdapterCache, error) {
	//if e.Redis.Addrs != nil {
	//	options, err := e.Redis.GetRedisOptions()
	//	if err != nil {
	//		return nil, err
	//	}
	//	r, err := cache.NewCusterRedis(GetCusterRedisClient(), e.Redis.Addrs, options)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return r, nil
	//}
	if e.Redis != nil {
		options, err := e.Redis.GetRedisOptions()
		if err != nil {
			return nil, err
		}
		r, err := cache.NewRedis(GetRedisClient(), options)
		if err != nil {
			return nil, err
		}
		if _redis == nil {
			_redis = r.GetClient()
		}
		return r, nil
	}
	return cache.NewMemory(), nil
}
