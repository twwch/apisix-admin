package apisix_sdk

import (
	"apisix-admin/proto/apisix/pb"
	"context"
	"fmt"
)

type Route struct {
	apisixClient
}

func (apisix *Route) List(ctx context.Context, page, size int) (resp *pb.ListRouteResp, err error) {
	path := fmt.Sprintf("/apisix/admin/routes?page=%d&size%d", page, size)
	err = apisix.client.Get(ctx, path, nil, &resp)
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