package redis_client

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RedisClient = ConnectRedis()

func ConnectRedis() *redis.Client {
	//TODO: Add your redis address
	store := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", DB: 0})
	fmt.Println("redis client connected")
	return store
}
