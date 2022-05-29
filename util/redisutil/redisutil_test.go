package redisutil_test

import (
	"context"
	"douyin/config"
	"douyin/util/redisutil"
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestRedisConn(t *testing.T) {
	ctx := context.Background()
	redisUtil := redisutil.NewRedisUtil(&redis.Pool{
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
	err := redisUtil.Set(ctx, "test", 123, 300000)
	if err != nil {
		t.Errorf("设置失败%v", err)
	}
	var value int
	hit, errget := redisUtil.Get(ctx, "test", &value)
	if errget != nil {
		t.Errorf("获取失败%v", errget)
	}
	if !hit {
		t.Error("key不存在")
	}
	fmt.Println(value)
}
