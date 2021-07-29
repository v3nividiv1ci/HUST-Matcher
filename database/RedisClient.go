package database

import "github.com/go-redis/redis"

var RClient *redis.Client

func InitRClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	RClient = client
	return client
}

func GetRClient() *redis.Client {
	return RClient
}

func Insertcode() {

}
