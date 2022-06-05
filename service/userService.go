package service

import (
	"context"
	"douyin/constants"
	"douyin/db"
	"douyin/response"
	"douyin/util/authutil"
	"douyin/util/md5util"
	"sync"
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

var (
	auth             *authutil.AuthUtil
	loadAuthUtilOnce sync.Once
)

func NewUserService() *UserSerice {
	loadAuthUtilOnce.Do(func() {
		auth = authutil.NewAuthUtil()
	})
	return &UserSerice{}
}

func (us *UserSerice) Login(username string, password string) (resp *UserLoginResponse) {
	userDao := db.GetUserByUsername(username)
	if userDao == nil {
		return &UserLoginResponse{
			Response: response.Response{
				StatusCode: 3001,
				StatusMsg:  constants.USER_NOT_EXIST_ERROR,
			},
		}
	}
	if md5util.MD5WithSalt(username, password) != userDao.Password {
		return &UserLoginResponse{
			Response: response.Response{
				StatusCode: 500,
				StatusMsg:  constants.PASSWORD_INCORRECT_ERROR,
			},
		}
	}
	token, err := auth.CreateToken(context.Background(), userDao.ID)
	if err != nil {
		return &UserLoginResponse{
			Response: response.Response{
				StatusCode: 500,
				StatusMsg:  "系统错误",
			},
		}
	}
	resp = &UserLoginResponse{
		Response: response.Response{
			StatusCode: 200,
			StatusMsg:  "登录成功",
		},
		UserId: int64(userDao.ID),
		Token:  token,
	}
	return resp
}
