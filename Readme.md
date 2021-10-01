
# docker api

文档： https://docs.docker.com/engine/api/v1.41/#operation/ContainerList

```shell script
修改配置文件 设置api端口
vi /lib/systemd/system/docker.service
修改ExecStart的值
[Service]
ExecStart=/usr/bin/dockerd -H unix://var/run/docker.sock -H tcp://0.0.0.0:2375

设置保存 重启docker

systemctl daemon-reload
systemctl restart docker

docker veriosn 可以看到对应的api版本

Docker version 20.10.7, build f0df350
[root@chtw ~]# docker version
Client: Docker Engine - Community
 Version:           20.10.7
 API version:       1.41
 Go version:        go1.13.15
 Git commit:        f0df350
 Built:             Wed Jun  2 11:58:10 2021
 OS/Arch:           linux/amd64
 Context:           default
 Experimental:      true

Server: Docker Engine - Community
 Engine:
  Version:          20.10.7
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.13.15
  Git commit:       b0f5bc3
  Built:            Wed Jun  2 11:56:35 2021
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.4.6
  GitCommit:        d71fcd7d8303cbf684402823e425e9dd2e99285d
 runc:
  Version:          1.0.0-rc95
  GitCommit:        b9ee9c6314599f1b4a7f497e1f1f856fe433d3b7
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0

```

```go
package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

//import "apisix-admin/cmd"

func main() {
	//cmd.Execute()
	ops := make([]client.Opt, 0)
	ops = append(ops, client.WithHost("tcp://127.0.0.1:2375"))
	ops = append(ops, client.WithVersion("v1.40"))
	cli, err := client.NewClientWithOpts(ops...)
	//cli, err :=  client.NewClient("tcp://192.168.209.152:2375", "v1.41", nil, nil)
	log(err)
	listImage(cli)
}
func listImage(cli *client.Client) {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	log(err)

	for _, image := range images {
		fmt.Printf("%+v\n",image)
	}
}

func log(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}
}
```

输出结果如下
```shell script
{Containers:-1 Created:1623743488 ID:sha256:5deb3699f2f70fdbddc3006c5bf1e3ec842bad610c19dd4710cf4fd07060fc54 Labels:map[] ParentID: RepoDigests:[apachepulsar/pulsar-all@sha256:bb497f4e036f180d18008f50ecea3fb29e69283a13690ad845ca0e6e67615da4] RepoTags:[apachepulsar/pulsar-all:latest] SharedSize:-1 Size:2905253524 VirtualSize:2905253524}
{Containers:-1 Created:1623742259 ID:sha256:8a8d4c7f511b1157d381ae44fbbd0eb578a2be7378eccef278437427d322c7d7 Labels:map[maintainer:Apache Pulsar <dev@pulsar.apache.organization>] ParentID: RepoDigests:[apachepulsar/pulsar-dashboard@sha256:804b48eab9ed94f11666f0009ab7ffab2297372b00410e2a0098fec73dc638aa] RepoTags:[apachepulsar/pulsar-dashboard:latest] SharedSize:-1 Size:1397399020 VirtualSize:1397399020}
{Containers:-1 Created:1537291564 ID:sha256:fd323a11901ef5c4d4d3f3adddba99a460cb22c78e2187a8eba46e7ca0b0d8f6 Labels:map[] ParentID: RepoDigests:[apachepulsar/pulsar@sha256:397b8915fa1e29e22af1a8805c34bab67b40ce1c6f44cc305ebda0deacf48ea4] RepoTags:[apachepulsar/pulsar:2.1.1-incubating] SharedSize:-1 Size:963831377 VirtualSize:963831377}
```