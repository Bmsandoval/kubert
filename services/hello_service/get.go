package hello_service

import (
	"context"
	"github.com/bmsandoval/kubert/grpc/protos"
	"time"
)

func (h *Helper) Get(request protos.HelloRequest) (response string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := h.AppCtx.Grpc.GreeterClient.SayHello(ctx, &request)
	if err != nil {
		return "", err
	}

	return r.GetMessage(), nil
}
