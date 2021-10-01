package apisix_sdk

import (
	"errors"
	"github.com/twwch/gin-sdk/httpclient/base"
)

var __client *apisixClient
var (
	URIORURLSChooseOneError = errors.New("URI和URIS不能同时使用")
)

func newErrors(msg string) error {
	return errors.New(msg)
}

func GetApiSixClient() *apisixClient {
	return __client
}

type apisixClient struct {
	client base.HttpClient
}

func (apisix *apisixClient) GetRoute() *Route {
	return &Route{apisixClient{client: apisix.client}}
}

func (apisix *apisixClient) GetUpstream() *Upstream {
	return &Upstream{apisixClient{client: apisix.client}}
}

func NewApiSixClient(host, key string) error {
	httpClient, err := base.NewClient(base.SetHost(host), base.SetHeaders(map[string]string{"X-API-KEY": key}))
	if err != nil {
		return err
	}
	__client = &apisixClient{
		client: httpClient,
	}
	return nil
}

