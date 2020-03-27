package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("init failed", err)
	}

	fmt.Println("init success")

	//zset

	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score: 90, Member: "java"},
		&redis.Z{Score: 95, Member: "php"},
		&redis.Z{Score: 96, Member: "c++"},
		&redis.Z{Score: 91, Member: "golang"},
	}
	_, err1 := redisDB.ZAdd(key, items...).Result()
	if err1 != nil {
		fmt.Println("zAdd failed")
		return
	}

	newScore, err2 := redisDB.ZIncrBy(key, 10.0, "golang").Result()
	if err2 != nil {
		fmt.Println("add failed")
		return
	}
	fmt.Println(newScore)

}

func initDB() (err error) {
	redisDB = redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               "127.0.0.1:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
	})

	_, err = redisDB.Ping().Result()
	return
}
