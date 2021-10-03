package route

import "apisix-admin/proto/apisix/pb"

type ListRoutes struct {
	Total  int32            `json:"total"`
	Routes []*pb.RouteNodes `json:"routes"`
}

type DeleteRouteReq struct {
	RouteId    string `json:"route_id" form:"route_id"`
	UpstreamId string `json:"upstream_id" form:"upstream_id"`
}

type CreateRouteReq struct {
	Route    *pb.CreateRouteReq    `json:"route" banding:"require"`
	Upstream *pb.CreateUpstreamReq `json:"upstream" banding:"require"`
}
