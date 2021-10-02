package apisix

import (
	apisix_sdk "apisix-admin/apisix-sdk"
	"apisix-admin/entity/apisix/route"
	"apisix-admin/proto/apisix/pb"
	"context"
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

func (app *ApisixApplication) DeleteRoute(ctx context.Context, routeId, upstreamId string) (err error) {
	_, err = apisix_sdk.GetApiSixClient().GetRoute().Delete(ctx, routeId)
	if err != nil{
		return
	}
	if upstreamId != "" {
		 _, err = apisix_sdk.GetApiSixClient().GetUpstream().Delete(ctx, upstreamId)
	}
	return
}
