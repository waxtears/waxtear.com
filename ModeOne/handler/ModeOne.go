package handler

import (
	"context"

	ModelOneProto "github.com/waxtears/waxtear.com/ModeOne/proto/ModelOneProto"
)

type ModeOne struct {
}

func (*ModeOne) FuncOne(ctx context.Context, req *ModelOneProto.FuncOneReq, rsp *ModelOneProto.FuncOneRsp) error {
	msg := "funcOne: " + req.GetName()
	rsp.Result = msg
	return nil
}
