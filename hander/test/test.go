package test

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/twwch/gin-sdk/handler"
	"github.com/twwch/gin-sdk/server"
	"net/http"
)

type TestHanlder struct {
	handler.Base
}

func (h *TestHanlder) Init(ginRouter *gin.RouterGroup) {
	h.Logger = log.WithField("handler", "test")
	// registry http handler
	if ginRouter != nil {
		appGroup := ginRouter.Group("/test")
		server.Route(appGroup, http.MethodGet, "/get", h.TestGet)
		server.Route(appGroup, http.MethodPost, "/post", h.TestPost)
		server.Route(appGroup, http.MethodPatch, "/patch", h.TestPatch)
		server.Route(appGroup, http.MethodPut, "/put", h.TestPut)
		server.Route(appGroup, http.MethodDelete, "/delete", h.TestDelete)
		//appGroup.GET("/", warper.CreateHandlerFunc(h.Test, false))
	}
}
type Req struct {
	Code int `json:"code" form:"code"`
}

type Resp struct {
	MyData map[string]interface{} `json:"my_data"`
}

func (h *TestHanlder) TestGet(ctx context.Context, req *Req) (*Resp, error) {
	h.Logger.Info(req)
	return &Resp{MyData:map[string]interface{}{"name": "Alice", "age": 12}}, nil
}

func (h *TestHanlder) TestPost(ctx context.Context, req *Req) (*Resp, error) {
	h.Logger.Info(req)
	return nil, errors.New("xxx")
}
func (h *TestHanlder) TestPatch(ctx context.Context, req *Req) (*Resp, error) {
	h.Logger.Info(req)
	return nil, errors.New("xxx")
}
func (h *TestHanlder) TestPut(ctx context.Context, req *Req) (*Resp, error) {
	h.Logger.Info(req)
	return nil, errors.New("xxx")
}
func (h *TestHanlder) TestDelete(ctx context.Context, req *Req) (*Resp, error) {
	h.Logger.Info(req)
	return nil, errors.New("xxx")
}

