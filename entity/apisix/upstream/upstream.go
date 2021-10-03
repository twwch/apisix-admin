package upstream

import "apisix-admin/proto/apisix/pb"

type ListUpstreamResp struct {
	Total     int32              `json:"total"`
	Upstreams []*pb.UpstreamNodes `json:"upstreams"`
}
