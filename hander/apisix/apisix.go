package apisix

import (
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
}

func (h *ApisixHanlder) Init(ginRouter *gin.RouterGroup) {
	h.Logger = log.WithField("handler", "ApisixHanlder")
	// registry http handler
	if ginRouter != nil {
		appGroup := ginRouter.Group("/apisix")

		// 路由相关
		routeGroup := appGroup.Group("/route")
		server.Route(routeGroup, http.MethodGet, "/get", h.ListRouter)
		//appGroup.GET("/", warper.CreateHandlerFunc(h.Test, false))
	}
}

func (h *ApisixHanlder) ListRouter(ctx context.Context, req *pb.ListReq) (resp *pb.ListRouteResp, err error) {
	return
}
