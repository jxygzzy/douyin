package service

import (
	"context"
	"douyin/db"
	"douyin/response"
	"douyin/util/authutil"
	"douyin/util/md5util"
	"douyin/util/randomutil"
	"fmt"
)

type UserSerice struct {
}

type UserLoginResponse struct {
	response.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	response.Response
	User response.User `json:"user"`
}

func NewUserService() *UserSerice {
	return &UserSerice{}
}

func (us *UserSerice) Login(username string, password string) (resp *UserLoginResponse, err error) {
	userDao, err := db.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if userDao == nil || md5util.MD5WithSalt(username, password) != userDao.Password {
		return nil, fmt.Errorf("用户名或密码不正确")
	}
	auth := authutil.NewAuthUtil()
	token, err := auth.CreateToken(context.Background(), userDao.ID)
	if err != nil {
		return nil, fmt.Errorf("系统缓存错误")
	}
	resp = &UserLoginResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "登录成功",
		},
		UserId: userDao.ID,
		Token:  token,
	}
	return resp, nil
}

func (us *UserSerice) Register(username string, password string) (*UserLoginResponse, error) {
	userDao, err := db.GetUserByUsername(username)
	if err == nil && userDao.ID != 0 {
		return nil, fmt.Errorf("用户名已存在")
	}
	err = nil
	userDao, err = db.Register(username, md5util.MD5WithSalt(username, password), "抖声用户"+randomutil.RandString(4))
	if err != nil {
		return nil, err
	}
	auth := authutil.NewAuthUtil()
	token, err := auth.CreateToken(context.Background(), userDao.ID)
	if err != nil {
		return nil, err
	}
	return &UserLoginResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "注册成功",
		},
		UserId: userDao.ID,
		Token:  token,
	}, nil
}

type UserInfoResponse struct {
	response.Response
	User *response.User `json:"user"`
}

func (us *UserSerice) UserInfo(user_id int64) (*UserInfoResponse, error) {
	user, err := db.GetUserById(user_id, user_id)
	if err != nil {
		return nil, err
	}
	return &UserInfoResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "获取成功",
		},
		User: user,
	}, nil
}
