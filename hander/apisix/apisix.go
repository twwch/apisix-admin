package apisix

import (
	"apisix-admin/application/apisix"
	"apisix-admin/entity/apisix/route"
	apisixEntity "apisix-admin/entity/apisix/route"
	"apisix-admin/entity/apisix/upstream"
	"apisix-admin/entity/common"
	"apisix-admin/proto/apisix/pb"
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/twwch/gin-sdk/handler"
	"github.com/twwch/gin-sdk/server"
	"net/http"
)

type ApisixHanlder struct {
	handler.Base
	*apisix.ApisixApplication
}

func (h *ApisixHanlder) Init(ginRouter *gin.RouterGroup) {
	h.Logger = log.WithField("handler", "ApisixHanlder")
	// registry http handler
	if ginRouter != nil {
		appGroup := ginRouter.Group("/apisix")

		// 路由相关
		routeGroup := appGroup.Group("/route")
		server.Route(routeGroup, http.MethodGet, "/list", h.ListRoute)
		server.Route(routeGroup, http.MethodGet, "/get", h.GetRoute)
		server.Route(routeGroup, http.MethodPost, "/create", h.CreateRoute)
		// 删除路由规则后， 服务无法访问， 谨慎操作
		//server.Route(routeGroup, http.MethodGet, "/delete", h.DeleteRoute)

		// Upstream相关
		UpstreamGroup := appGroup.Group("/upstream")
		server.Route(UpstreamGroup, http.MethodGet, "/list", h.ListUpstream)
		//appGroup.GET("/", warper.CreateHandlerFunc(h.Test, false))
	}
}

func (h *ApisixHanlder) ListRoute(ctx context.Context, req *pb.ListReq) (resp *route.ListRoutes, err error) {
	// apisix 分页无效，这个参数可以改为空
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 20
	}
	resp, err = h.ApisixApplication.ListRoute(ctx, req)
	if err != nil {
		h.Logger.Error(err)
	}
	return
}

func (h *ApisixHanlder) CreateRoute(ctx context.Context, req *apisixEntity.CreateRouteReq) (resp *common.Empty, err error) {
	if req.Upstream == nil || req.Route == nil {
		err = common.ParamsError
		return
	}
	resp, err = h.ApisixApplication.CreateRoute(ctx, req)
	if err != nil {
		h.Logger.Error(err)
	}
	return
}

func (h *ApisixHanlder) GetRoute(ctx context.Context, req *apisixEntity.GetRouteReq) (resp *apisixEntity.GetRouteResp, err error) {
	if req.Id == "" {
		err = common.ParamsError
		return
	}
	resp, err = h.ApisixApplication.GetRoute(ctx, req)
	if err != nil {
		h.Logger.Error(err)
	}
	return
}

func (h *ApisixHanlder) DeleteRoute(ctx context.Context, req *route.DeleteRouteReq) (resp *common.Empty, err error) {
	if req.RouteId == "" {
		err = common.ParamsError
		return
	}
	err = h.ApisixApplication.DeleteRoute(ctx, req.RouteId, req.UpstreamId)
	if err != nil {
		h.Logger.Error(err)
	}
	return
}

func (h *ApisixHanlder) ListUpstream(ctx context.Context, req *pb.ListReq) (resp *upstream.ListUpstreamResp, err error) {
	resp, err = h.ApisixApplication.ListUpstream(ctx, req)
	if err != nil {
		h.Logger.Error(err)
	}
	return
}
