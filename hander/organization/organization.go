package organization

import (
	"apisix-admin/application/organization"
	"apisix-admin/entity/common"
	"apisix-admin/entity/organization/user"
	"apisix-admin/entity/organization/user/mongo"
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/twwch/gin-sdk/handler"
	"github.com/twwch/gin-sdk/server"
	"net/http"
)

type OrganizationHanlder struct {
	handler.Base
	*organization.OrganizationApplication
}

func (h *OrganizationHanlder) Init(ginRouter *gin.RouterGroup) {
	h.Logger = log.WithField("handler", "OrganizationHanlder")
	h.OrganizationApplication = organization.NewOrganizationApplication()
	// registry http handler
	if ginRouter != nil {
		appGroup := ginRouter.Group("/organization")

		// 用户相关
		userGroup := appGroup.Group("/user")
		server.Route(userGroup, http.MethodPost, "/login", h.LoginHandler)
		server.Route(userGroup, http.MethodGet, "/info", h.InfoHandler)

	}
}

func (h *OrganizationHanlder) LoginHandler(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = h.OrganizationApplication.Login(ctx, req)
	if err != nil{
		h.Logger.Error(err)
	}
	return
}

func (h *OrganizationHanlder) InfoHandler(ctx context.Context, req *common.Empty) (resp *mongo.User, err error)  {
	resp, err = h.OrganizationApplication.Info(ctx)
	if err != nil{
		h.Logger.Error(err)
		err = common.SystemError
	}
	return
}