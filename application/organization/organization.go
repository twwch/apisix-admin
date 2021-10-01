package organization

import (
	"apisix-admin/domain/user"
	userEntity "apisix-admin/entity/user"
	"apisix-admin/proto/apisix/pb"
	"context"
)

type OrganizationApplication struct {
	userService *user.UserService
}

func NewOrganizationApplication() *OrganizationApplication {
	return &OrganizationApplication{userService:user.NewUserService()}
}

func (app *OrganizationApplication) Login(ctx context.Context, req *pb.ListReq)(resp *userEntity.User, err error) {
	return app.userService.FindUserByIdOrEmail(ctx, "chenhao03", false)
}
