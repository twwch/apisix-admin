package apisix_sdk

import (
	"apisix-admin/proto/apisix/pb"
	"context"
	"fmt"
	"log"
	"testing"
)

func TestUpstream_List(t *testing.T) {
	Init()
	ctx := context.Background()
	resp, err := GetApiSixClient().GetUpstream().List(ctx, 2, 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Action, resp.Count, resp.Node )
	for _, item := range resp.Node.Nodes{
		GetApiSixClient().GetUpstream().Delete(ctx, item.Value.Id)
	}
}

func TestUpstream_Create(t *testing.T) {
	Init()
	ctx := context.Background()
	resp, err := GetApiSixClient().GetUpstream().Create(ctx, &pb.CreateUpstreamReq{
		Name:"测试",
		Desc:"cesi",
		Nodes: map[string]int32{"127.0.0.1:80":1, "127.0.0.1:81":1},
		Type: pb.UpstreamRoundrobinType,
		Id: "4",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Node, resp.Action, resp.Node)
}

func TestUpstream_Detele(t *testing.T) {
	Init()
	ctx := context.Background()
	resp, err := GetApiSixClient().GetUpstream().Delete(ctx,  "1633250139")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Deleted, resp.Action)
}