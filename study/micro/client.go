package main

import (
	"context"
	"fmt"
	"micro/TestMicroProto"

	"github.com/asim/go-micro/v3"
)

func main() {
	server := micro.NewService(micro.Name("waxtear.com/test/micro.client"))

	server.Init()

	client := TestMicroProto.NewTestService("waxtear.com/test/micro.server", server.Client())

	res, err := client.SayHello(context.TODO(), &TestMicroProto.SayHelloReq{Name: "123"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
}
