package apisix

import (
	apisix_sdk "apisix-admin/apisix-sdk"
	"apisix-admin/entity/apisix/route"
	apisixEntity "apisix-admin/entity/apisix/route"
	"apisix-admin/entity/apisix/upstream"
	"apisix-admin/entity/common"
	"apisix-admin/proto/apisix/pb"
	"context"
	"fmt"
	"time"
)

type ApisixApplication struct {
}

func (app *ApisixApplication) ListRoute(ctx context.Context, req *pb.ListReq) (resp *route.ListRoutes, err error) {
	routes, err := apisix_sdk.GetApiSixClient().GetRoute().List(ctx, req.Page, req.Size)
	if err != nil {
		return
	}
	resp = new(route.ListRoutes)
	resp.Total = routes.Count
	resp.Routes = routes.Node.Nodes
	return
}

func (app *ApisixApplication) CreateRoute(ctx context.Context, req *apisixEntity.CreateRouteReq) (resp *common.Empty, err error) {
	id := fmt.Sprint(time.Now().Unix())
	req.Upstream.Id = id
	req.Upstream.Type = pb.UpstreamRoundrobinType
	_, err = apisix_sdk.GetApiSixClient().GetUpstream().Create(ctx, req.Upstream)
	if err != nil {
		return
	}
	req.Route.Id = id
	req.Route.UpstreamId = id
	_, err = apisix_sdk.GetApiSixClient().GetRoute().Create(ctx, req.Route)
	return
}

func (app *ApisixApplication) DeleteRoute(ctx context.Context, routeId, upstreamId string) (err error) {
	_, err = apisix_sdk.GetApiSixClient().GetRoute().Delete(ctx, routeId)
	if err != nil {
		return
	}
	if upstreamId != "" {
		_, err = apisix_sdk.GetApiSixClient().GetUpstream().Delete(ctx, upstreamId)
	}
	return
}

func (app *ApisixApplication) ListUpstream(ctx context.Context, req *pb.ListReq) (resp *upstream.ListUpstreamResp, err error) {
	upstreamItem, err := apisix_sdk.GetApiSixClient().GetUpstream().List(ctx, req.Page, req.Size)
	if err != nil {
		return
	}
	resp = new(upstream.ListUpstreamResp)
	resp.Total = upstreamItem.Count
	resp.Upstreams = upstreamItem.Node.Nodes
	return
}
