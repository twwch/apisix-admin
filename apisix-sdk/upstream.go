package apisix_sdk

import (
	"apisix-admin/proto/apisix/pb"
	"context"
	"fmt"
)

type Upstream struct {
	apisixClient
}

func (upstream *Upstream) List(ctx context.Context, page, size int) (resp *pb.ListUpstreamResp, err error) {
	path := fmt.Sprintf("/apisix/admin/upstreams?page=%d&&size=%d", page, size)
	err = upstream.client.Get(ctx, path, nil, &resp)
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

func (upstream *Upstream) Detele(ctx context.Context, id string) (resp *pb.DeleteResp, err error) {
	path := fmt.Sprintf("/apisix/admin/upstreams/%s", id)
	err = upstream.client.Delete(ctx, path, nil, &resp)
	return
}
