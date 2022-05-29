package authutil

import (
	"context"
	"douyin/config"
	"douyin/util/redisutil"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type AuthUtil struct {
	ru *redisutil.RedisUtil
}

var (
	loadRedisUtilOnce sync.Once
)

func (a *AuthUtil) loadRedisUtil() {
	// 单例模式
	loadRedisUtilOnce.Do(func() {
		a.ru = redisutil.NewRedisUtil(&redis.Pool{
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", config.Redis_addr,
					redis.DialDatabase(config.Redis_db),
					redis.DialPassword(config.Redis_password))
				if err != nil {
					return nil, err
				}
				return c, nil
			},
		})
	})

}

func NewAuthUtil() *AuthUtil {
	return &AuthUtil{}
}

func (a *AuthUtil) CreateToken(ctx context.Context, userId int) (string, error) {
	if a.ru == nil {
		a.loadRedisUtil()
	}
	token := uuid.New().String()
	err := a.ru.Set(ctx, token, userId, config.Redis_ttl)
	if err != nil {
		return "", err
	}
	return token, nil
}
