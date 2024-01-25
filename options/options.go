package options

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
)

type clientSuiteFn func() []client.Option

type serverSuiteFn func() []server.Option

func (f clientSuiteFn) Options() []client.Option {
	return f()
}

func (f serverSuiteFn) Options() []server.Option {
	return f()
}

func ClientSuite(opts ...client.Option) client.Suite {
	return clientSuiteFn(func() []client.Option {
		return opts
	})
}

func ServerSuite(opts ...server.Option) server.Suite {
	return serverSuiteFn(func() []server.Option {
		return opts
	})
}
