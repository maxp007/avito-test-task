package cache

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/maxp007/avito-test-task/config"
	_ "github.com/maxp007/avito-test-task/config"
	"log"
)

var Cache *redis.Client

func init() {
	host := config.GetInstance().Data.Cache.Host
	port := config.GetInstance().Data.Cache.Port

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
		return
	}
	Cache = client
	return
}

func ConnClose() (err error) {
	defer func(err *error) {
		(*err) = Cache.Close()

	}(&err)
	return err
}
