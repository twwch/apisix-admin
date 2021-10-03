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
	if req.Id != ""{
		id = req.Id
	}
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

func (app *ApisixApplication) GetRoute(ctx context.Context, req *apisixEntity.GetRouteReq) (resp *apisixEntity.GetRouteResp, err error) {
	routes, err := apisix_sdk.GetApiSixClient().GetRoute().Get(ctx, req.Id)
	if err != nil {
		return
	}
	upstreams, err := apisix_sdk.GetApiSixClient().GetUpstream().Get(ctx, req.Id)
	if err != nil {
		return
	}
	resp = new(apisixEntity.GetRouteResp)
	routeItem := routes.Node.Value
	resp.Route = &pb.CreateRouteReq{
		Id:          routeItem.Id,
		Uris:        routeItem.Uris,
		Uri:         routeItem.Uri,
		Hosts:       routeItem.Hosts,
		Desc:        routeItem.Desc,
		RemoteAddrs: routeItem.RemoteAddrs,
		UpstreamId:  routeItem.UpstreamId,
		Methods:     routeItem.Methods,
	}
	upstreamItem := upstreams.Node.Value
	resp.Upstream = &pb.CreateUpstreamReq{
		Id: upstreamItem.Id,
		Nodes: upstreamItem.Nodes,
		Name: upstreamItem.Name,
		Desc: upstreamItem.Desc,
		Type: upstreamItem.Type,
	}
	nodes := make([]*apisixEntity.Nodes, 0)
	for key, item := range upstreams.Node.Value.Nodes {
		nodes = append(nodes, &apisixEntity.Nodes{
			NodeKey:   key,
			NodeValue: item,
		})
	}
	resp.Nodes = nodes
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
