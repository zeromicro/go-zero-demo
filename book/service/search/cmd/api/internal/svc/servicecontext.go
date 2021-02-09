package svc

import (
	"go-zero-demo/book/service/search/cmd/api/internal/config"
	"go-zero-demo/book/service/search/cmd/api/internal/middleware"
	"go-zero-demo/book/service/user/cmd/rpc/userclient"

	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Example rest.Middleware
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Example: middleware.NewExampleMiddleware().Handle,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
