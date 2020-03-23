package appcontext

import (
	"context"
	"github.com/bmsandoval/kubert/configs"
	"github.com/bmsandoval/kubert/services/grpc_service"
)

type Context struct {
	Config    configs.Configuration
	GoContext context.Context
	Grpc      grpc_service.Connection
}