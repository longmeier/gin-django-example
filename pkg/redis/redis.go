package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var Client *redis.Client

func NewRedis(conf *viper.Viper) {
	addr := fmt.Sprintf("%v:%v", conf.GetString("redis.host"), conf.GetString("redis.port"))
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: conf.GetString("redis.pwd"),
		DB:       conf.GetInt("redis.db"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("redis error: %s", err.Error()))
	}

}
