package organization

import (
	"apisix-admin/domain/user"
	entityUser "apisix-admin/entity/organization/user"
	"apisix-admin/entity/organization/user/mongo"
	"apisix-admin/utils/des"
	"apisix-admin/utils/jwt"
	"context"
	"errors"
	"github.com/twwch/gin-sdk/middles"
	"net/http"
)

type OrganizationApplication struct {
	userService *user.UserService
}

func NewOrganizationApplication() *OrganizationApplication {
	return &OrganizationApplication{userService:user.NewUserService()}
}

func (app *OrganizationApplication) Login(ctx context.Context, req *entityUser.LoginReq) (resp *entityUser.LoginResp, err error) {
	pass := des.GetPass(req.Password)
	if pass == ""{
		err = entityUser.LoginFailError
		return
	}
	userTemp , err := app.userService.FindUserByPhoneOrEmail(ctx, req.Account, pass)
	if err != nil{
		err = errors.New("获取用户失败，用户账号或密码错误")
		return
	}
	token, err := jwt.MakeToken(userTemp.UserId)
	if err != nil{
		return
	}
	resp = new(entityUser.LoginResp)
	resp.UserInfo = userTemp
	resp.Token = token
	return
}

func(app *OrganizationApplication)  Info(ctx context.Context) (resp *mongo.User, err error) {
	req := ctx.Value(middles.RequestKey).(*http.Request)
	token := req.Header.Get("token")
	_, claims, err := jwt.ParseToken(token)
	if err != nil{
		return nil, err
	}
	resp, err = app.userService.FindUserByUserId(ctx, claims.UserId)
	return
}