package apisix_sdk

import (
	"apisix-admin/proto/apisix/pb"
	"context"
	"fmt"
	"strings"
)

type Route struct {
	apisixClient
}

func (apisix *Route) List(ctx context.Context, page, size int32) (resp *pb.ListRouteResp, err error) {
	// apisix 分页无效，page， size 参数可以改为空
	path := fmt.Sprintf("/apisix/admin/routes?page=%d&size=%d", page, size)
	err = apisix.client.Get(ctx, path, nil, &resp)
	// 这是一个坑，当没有路由存在的时候，接口返回的{}， 有数据的时候返回的是数组
	if err != nil && strings.Contains(err.Error(), "cannot unmarshal object into Go struct field RouteNode.node.nodes"){
		err = nil
	}
	return
}

func (apisix *Route) Get(ctx context.Context, id string) (resp *pb.GetRouteResp, err error) {
	// apisix 分页无效，page， size 参数可以改为空
	path := fmt.Sprintf("/apisix/admin/routes/%s", id)
	err = apisix.client.Get(ctx, path, nil, &resp)
	// 这是一个坑，当没有路由存在的时候，接口返回的{}， 有数据的时候返回的是数组
	if err != nil && strings.Contains(err.Error(), "cannot unmarshal object into Go struct field RouteNode.node.nodes"){
		err = nil
	}
	return
}

func (apisix *Route) Create(ctx context.Context, req *pb.CreateRouteReq) (resp *pb.CreateRouteResp, err error) {
	if req.GetUri() != "" && len(req.GetUris()) > 0 {
		err = URIORURLSChooseOneError
		return
	}
	path := fmt.Sprintf("/apisix/admin/routes/%s", req.GetId())
	params := make(map[string]interface{})
	if len(req.GetUris()) > 0 {
		params["uris"] = req.GetUris()
	}
	if req.Uri != "" {
		params["uri"] = req.GetUri()
	}
	if req.GetDesc() != "" {
		params["desc"] = req.GetDesc()
	}
	if len(req.GetRemoteAddrs()) > 0 {
		params["remote_addrs"] = req.GetRemoteAddrs()
	}
	if len(req.GetHosts()) > 0 {
		params["hosts"] = req.GetHosts()
	}
	if len(req.GetMethods()) > 0 {
		params["methods"] = req.GetMethods()
	}
	if req.GetUpstreamId() != "" {
		params["upstream_id"] = req.GetUpstreamId()
	}
	err = apisix.client.Put(ctx, path, params, &resp)
	if err != nil {
		return
	}
	if resp.ErrorMsg != "" {
		err = newErrors(resp.ErrorMsg)
	}
	return
}

func (apisix *Route) Delete(ctx context.Context, id string) (resp *pb.DeleteResp, err error)  {
	path := fmt.Sprintf("/apisix/admin/routes/%s", id)
	err = apisix.client.Delete(ctx, path, nil, &resp)
	return
}