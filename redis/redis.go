package redis

import (
	"log"
	"time"
	"github.com/garyburd/redigo/redis"
	"github.com/shoother/g"
)

func InitRedis(redisConfig *g.RedisConfig) *redis.Pool {
	pwd := redisConfig.Password

	ConnPool := &redis.Pool{
		MaxIdle:     redisConfig.Idle,
		IdleTimeout: time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisConfig.Addr)
			if err != nil {
				return nil, err
			}

			if pwd != "" {
				if _, err := c.Do("AUTH", pwd); err != nil {
					c.Close()
					log.Println("[ERROR] redis auth fail")
					return nil, err
				}
			}

			return c, err
		},
		TestOnBorrow: PingRedis,
	}
	return ConnPool
}

func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
	}
	return err
}
