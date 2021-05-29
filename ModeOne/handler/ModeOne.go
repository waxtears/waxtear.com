package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	ModeOne "ModeOne/proto"
)

type ModeOne struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *ModeOne) Call(ctx context.Context, req *ModeOne.Request, rsp *ModeOne.Response) error {
	log.Info("Received ModeOne.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *ModeOne) Stream(ctx context.Context, req *ModeOne.StreamingRequest, stream ModeOne.ModeOne_StreamStream) error {
	log.Infof("Received ModeOne.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&ModeOne.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *ModeOne) PingPong(ctx context.Context, stream ModeOne.ModeOne_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&ModeOne.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
