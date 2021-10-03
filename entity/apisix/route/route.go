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
	Id       string                `json:"id"`
}

type GetRouteReq struct {
	Id string `json:"id" form:"id"`
}

type Nodes struct {
	NodeValue int32  `json:"node_value"`
	NodeKey   string `json:"node_key"`
}

type GetRouteResp struct {
	Route    *pb.CreateRouteReq    `json:"route"`
	Upstream *pb.CreateUpstreamReq `json:"upstream"`
	Nodes    []*Nodes              `json:"nodes"`
}
