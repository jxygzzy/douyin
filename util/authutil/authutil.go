package authutil

import (
	"context"
	"douyin/config"
	"douyin/constants"
	"douyin/util/redisutil"
	"errors"
	"log"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type AuthUtil struct {
}

var (
	redis_pool_config = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Redis_addr,
				redis.DialDatabase(config.Redis_db),
				redis.DialPassword(config.Redis_password))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
)

func NewAuthUtil() *AuthUtil {
	return &AuthUtil{}
}

func (a *AuthUtil) CreateToken(ctx context.Context, userId int64) (string, error) {
	ru := redisutil.NewRedisUtil(redis_pool_config)
	token := uuid.New().String()
	err := ru.Set(ctx, token, userId, config.Redis_ttl)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *AuthUtil) CheckToken(ctx context.Context, token string) (int64, error) {
	ru := redisutil.NewRedisUtil(redis_pool_config)
	var userId int64
	hit, err := ru.Get(ctx, token, &userId)
	if err != nil {
		return 0, err
	}
	if !hit {
		return 0, errors.New(constants.TOKEN_NOT_EXIST_ERROR)
	}
	go a.RefreshToken(ctx, token)
	return userId, nil
}

func (a *AuthUtil) RefreshToken(ctx context.Context, token string) {
	ru := redisutil.NewRedisUtil(redis_pool_config)
	ttl, ttlErr := ru.TTL(ctx, token)
	if ttlErr != nil {
		log.Printf("error when ttl token:%v", ttlErr)
	}
	if ttl < config.Redis_ttl-config.Redis_refresh {
		err := ru.Expire(ctx, token, config.Redis_ttl)
		if err != nil {
			log.Printf("error when reflash token:%v", err)
		}
	}
}
