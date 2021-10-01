package images

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ListImages(cli *client.Client) {
	images, _ := cli.ImageList(context.Background(), types.ImageListOptions{})

	for _, image := range images {
		fmt.Printf("%+v\n",image)
	}
}
