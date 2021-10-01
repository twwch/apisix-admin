package organization

import (
	"apisix-admin/application/organization"
	userEntity "apisix-admin/entity/user"
	"apisix-admin/proto/apisix/pb"
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
		server.Route(userGroup, http.MethodGet, "/login", h.LoginHandler)
	}
}

func (h *OrganizationHanlder) LoginHandler(ctx context.Context, req *pb.ListReq) (resp *userEntity.User, err error) {
	resp, err = h.OrganizationApplication.Login(ctx, req)
	if err != nil{
		h.Logger.Error(err)
	}
	return
}
