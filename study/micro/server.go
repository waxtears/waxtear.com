package main

import (
	"context"
	"fmt"

	micro_pb "micro/TestMicroProto"

	"github.com/asim/go-micro/v3"
)

type MicroServer struct{}

func (*MicroServer) SayHello(ctx context.Context, req *micro_pb.SayHelloReq, rsp *micro_pb.SayHelloRsp) error {
	rsp.Msg = "Hello " + req.GetName()
	return nil
}

func main() {
	server := micro.NewService(micro.Name("waxtear.com/test/micro.server"))
	server.Init()

	micro_pb.RegisterTestServiceHandler(server.Server(), new(MicroServer))

	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
