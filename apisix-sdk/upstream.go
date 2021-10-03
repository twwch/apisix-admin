package apisix_sdk

import (
	"apisix-admin/proto/apisix/pb"
	"context"
	"fmt"
	"strings"
)

type Upstream struct {
	apisixClient
}

func (upstream *Upstream) List(ctx context.Context, page, size int32) (resp *pb.ListUpstreamResp, err error) {
	path := fmt.Sprintf("/apisix/admin/upstreams?page=%d&&size=%d", page, size)
	err = upstream.client.Get(ctx, path, nil, &resp)
	// 这是一个坑，当没有路由存在的时候，接口返回的{}， 有数据的时候返回的是数组
	if err != nil && strings.Contains(err.Error(), "json: cannot unmarshal object into Go struct field UpstreamNode.node.nodes of type []*pb.UpstreamNodes"){
		err = nil
	}
	return
}

func (apisix *Upstream) Get(ctx context.Context, id string) (resp *pb.GetUpstreamResp, err error) {
	// apisix 分页无效，page， size 参数可以改为空
	path := fmt.Sprintf("/apisix/admin/upstreams/%s", id)
	err = apisix.client.Get(ctx, path, nil, &resp)
	// 这是一个坑，当没有路由存在的时候，接口返回的{}， 有数据的时候返回的是数组
	if err != nil && strings.Contains(err.Error(), "json: cannot unmarshal object into Go struct field UpstreamNode.node.nodes of type []*pb.UpstreamNodes"){
		err = nil
	}
	return
}

func (upstream *Upstream) Create(ctx context.Context, req *pb.CreateUpstreamReq) (resp *pb.CreateUpstreamResp, err error) {
	path := fmt.Sprintf("/apisix/admin/upstreams/%s", req.GetId())
	params := make(map[string]interface{})
	if req.GetDesc() != "" {
		params["desc"] = req.GetDesc()
	}
	if req.GetName() != "" {
		params["name"] = req.GetName()
	}
	if req.GetNodes() != nil {
		params["nodes"] = req.GetNodes()
	}
	if req.GetType() != "" {
		params["type"] = req.GetType()
	}
	err = upstream.client.Put(ctx, path, params, &resp)
	return
}

func (upstream *Upstream) Delete(ctx context.Context, id string) (resp *pb.DeleteResp, err error) {
	path := fmt.Sprintf("/apisix/admin/upstreams/%s", id)
	err = upstream.client.Delete(ctx, path, nil, &resp)
	return
}
