package user

import (
	"apisix-admin/entity/organization/user/mongo"
	"errors"
)

var LoginFailError = errors.New("登录失败")

type LoginReq struct {
	Account  string `json:"account" form:"account"`
	Password string `json:"password" form:"password"`
}

type LoginResp struct {
	Token    string      `json:"token"`
	UserInfo *mongo.User `json:"user_info"`
}
