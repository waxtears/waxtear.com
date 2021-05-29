package main

import (
	"ModeOne/handler"
	pb "ModeOne/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("modeone"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterModeOneHandler(srv.Server(), new(handler.ModeOne))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
