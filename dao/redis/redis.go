package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化Redis连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.auth"),   // no password set
		DB:       viper.GetInt("redis.db"),        // use default DB
		PoolSize: viper.GetInt("redis.pool_size"), // connection pool size
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		zap.L().Error("connect redis failed", zap.Error(err))
		return
	}
	return
}

func Close() {
	_ = rdb.Close()
}
