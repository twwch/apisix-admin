package server

import (
	"apisix-admin/config"
	"apisix-admin/hander/apisix"
	"apisix-admin/hander/organization"
	"apisix-admin/hander/test"
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/twwch/gin-sdk/handler"
	"github.com/twwch/gin-sdk/server"
)

// HTTP stands for HTTP server.
type HTTP struct {
	Server server.Server
}

var defaultHTTP Server = &HTTP{}

// New creates a new HTTP server.
func (h *HTTP) New(address string) Server {
	return &HTTP{
		Server: newRouter(address),
	}
}

var httphandlers = []handler.Handler{
	&test.TestHanlder{},
	&apisix.ApisixHanlder{},
	&organization.OrganizationHanlder{},
}

func newRouter(address string) server.Server {
	ctx := context.Background()
	conf := config.Get()
	httpServer := server.NewServer(server.Options{
		Name:    "apisix_admin",
		Address: address,
		LogConf: conf.Log,
	})

	router := httpServer.GetEngine()
	v1 := router.Group("/apisix_admin/v1")
	{
		for _, hd := range httphandlers {
			hd.Init(v1)
		}
	}
	log.Info(httpServer.Run(ctx))
	return httpServer
}
