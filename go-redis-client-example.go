package main

import (
	"fmt"

	//"github.com/garyburd/redigo/redis"
	"github.com/go-redis/redis" //redis的包这个最好用
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	if _, err := client.Ping().Result(); err != nil {
		fmt.Println(err)
	}
	return client
}

func MGetFeatures(keys []string) ([]interface{}, error) {
	redisClient := NewRedis()
	return redisClient.MGet(keys...).Result()
}

func main() {
	keys := []string{"key_1", "key_2"}
	fmt.Println(MGetFeatures(keys))
}
