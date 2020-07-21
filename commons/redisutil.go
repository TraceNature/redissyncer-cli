package commons

import (
	redis "github.com/go-redis/redis/v7"
)

//GetGoRedisClient 获取redis client
func GetGoRedisClient(opt *redis.Options) *redis.Client {
	client := redis.NewClient(opt)
	return client
}

func GetGoRedisConn(opt *redis.Options) *redis.Conn {
	client := redis.NewClient(opt)
	return client.Conn()
}

//redisserver联通性校验
func CheckRedisClientConnect(r *redis.Client) bool {
	_, err := r.Ping().Result()
	if err != nil {
		return false
	}
	return true
}

func CheckRedisClusterClientConnect(r *redis.ClusterClient) bool {
	_, err := r.Ping().Result()
	if err != nil {

		return false
	}
	return true
}
