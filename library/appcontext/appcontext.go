package appcontext

import (
	"context"
	"github.com/bmsandoval/kubert/configs"
	"github.com/bmsandoval/kubert/grpc"
)

type Context struct {
	Config configs.Configuration
	GoContext context.Context
	Grpc grpc.Connection
}