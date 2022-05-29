package authutil_test

import (
	"context"
	"douyin/util/authutil"
	"fmt"
	"testing"
)

func TestConnRedis(t *testing.T) {
	auth := authutil.NewAuthUtil()
	token, err := auth.CreateToken(context.Background(), 123)
	if err != nil {
		t.Errorf("创建失败%v", err)
	} else {
		fmt.Printf("token:%v", token)
	}

}
