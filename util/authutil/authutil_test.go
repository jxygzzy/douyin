package authutil_test

import (
	"context"
	"douyin/util/authutil"
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	Auth := authutil.NewAuthUtil()
	token, err := Auth.CreateToken(context.Background(), 123)
	if err != nil {
		t.Errorf("创建失败%v", err)
	} else {
		fmt.Printf("token:%v", token)
	}
}

func TestGetToken(t *testing.T) {
	Auth := authutil.NewAuthUtil()
	userId, err := Auth.CheckToken(context.Background(), "4393b8db-e007-4227-a72b-9c4d8b9d7d75")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("userId:%v", userId)
	}
}

func TestRefresh(t *testing.T) {
	Auth := authutil.NewAuthUtil()
	Auth.RefreshToken(context.Background(), "d1b0a88d-f94a-4255-8d5d-df7ed2b87723")
}
