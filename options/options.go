//go:build go1.19
// +build go1.19

package options

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
)

type options[T client.Option | server.Option] []T

func (opts options[T]) Options() []T {
	return opts
}

func ClientSuite(opts ...client.Option) client.Suite {
	return options[client.Option](opts)
}

func ServerSuite(opts ...server.Option) server.Suite {
	return options[server.Option](opts)
}
