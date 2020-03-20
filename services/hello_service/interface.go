package hello_service

import (
	"github.com/bmsandoval/kubert/grpc/protos"
	"github.com/bmsandoval/kubert/library/appcontext"
)

type Helper struct {
	AppCtx appcontext.Context
}
type Helpable struct{}

func(h Helpable) NewHelper(appCtx appcontext.Context) (interface{}, error) {
	return &Helper{
		AppCtx: appCtx,
	}, nil
}

func (h Helpable) ServiceName() string {
	return "HelloSvc"
}

type Service interface {
	Get(request protos.HelloRequest) (response string, err error)
}
