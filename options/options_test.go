package options_test

import (
	"testing"

	"github.com/bitbus/kx/options"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/server"
)

func TestClientSuite(t *testing.T) {
	suite := options.ClientSuite(
		client.WithStatsLevel(stats.LevelBase),
		client.WithMuxConnection(1),
		client.WithConnReporterEnabled(),
	)
	opts := suite.Options()
	if len(opts) != 3 {
		t.Errorf("want 3 options but got %d", len(opts))
	}
}

func TestServerSuite(t *testing.T) {
	suite := options.ServerSuite(
		server.WithMuxTransport(),
		server.WithContextBackup(true, true),
		server.WithReusePort(true),
	)
	opts := suite.Options()
	if len(opts) != 3 {
		t.Errorf("want 3 options but got %d", len(opts))
	}
}
