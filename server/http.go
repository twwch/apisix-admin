package server

import (
	"apisix-admin/config"
	"apisix-admin/hander/apisix"
	"apisix-admin/hander/organization"
	"apisix-admin/hander/test"
	"apisix-admin/middleware/jwt"
	"context"
	"github.com/gin-gonic/gin"
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
	v1.Use(CORSMiddleware(), jwt.JwtMiddleware())
	{
		for _, hd := range httphandlers {
			hd.Init(v1)
		}
	}
	log.Info(httpServer.Run(ctx))
	return httpServer
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
