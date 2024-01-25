package options

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
)

type clientOptions []client.Option

type serverOptions []server.Option

func (opts clientOptions) Options() []client.Option {
	return opts
}

func (opts serverOptions) Options() []server.Option {
	return opts
}

func ClientSuite(opts ...client.Option) client.Suite {
	return clientOptions(opts)
}

func ServerSuite(opts ...server.Option) server.Suite {
	return serverOptions(opts)
}
