package main

import (
	"fmt"

	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/waxtears/waxtear.com/ModeOne/handler"
	"github.com/waxtears/waxtear.com/ModeOne/proto/ModelOneProto"
)

func main() {
	// Create service
	svr := micro.NewService(
		micro.Name("modeone"),
		micro.Version("latest"),
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
