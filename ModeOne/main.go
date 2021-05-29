package main

import (
	"fmt"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"

	//"github.com/asim/nitro/v3/registry"
	//"github.com/asim/nitro-plugins/registry/consul/v3"
	"github.com/waxtears/waxtear.com/ModeOne/handler"
	"github.com/waxtears/waxtear.com/ModeOne/proto/ModelOneProto"
)

func main() {
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create service
	svr := micro.NewService(
		micro.Name("modeone"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulRegistry),
	)

	svr.Init()

	// Register handler
	err := ModelOneProto.RegisterModelOneProtoHandler(svr.Server(), new(handler.ModeOne))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run service
	if err := svr.Run(); err != nil {
		logger.Fatal(err)
		return
	}

	return
}
