package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-talk-talk/config"
	"go-talk-talk/global"
	"time"
)

func init() {
	// 新版放弃了原有的NewPool方法，可以通过直接初始化Pool来获得一个连接池
	global.RedisPool = &redis.Pool{
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			fmt.Println(time.Since(t))
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return redis.Dial(config.RedisConf.Network, fmt.Sprintf("%s:%s", config.RedisConf.Host, config.RedisConf.Port))
		},
	}
}
func Get() redis.Conn {
	return global.RedisPool.Get()
}
