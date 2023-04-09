package rpcgw

import (
	"context"
	"net/http"
	"time"

	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const rpcClientTimeout = 4 * time.Second

type Gateway struct {
	ctx    context.Context
	url    string
	client jsonrpc.RPCClient
	logger log.Logger
}

func NewGateway(logger log.Logger, url string) ports.RPCGateway {
	httpClient := &http.Client{
		Timeout: rpcClientTimeout,
	}

	opts := &jsonrpc.RPCClientOpts{
		HTTPClient: httpClient,
	}

	rpcClient := jsonrpc.NewClientWithOpts(url, opts)

	return &Gateway{
		ctx:    context.Background(),
		url:    url,
		client: rpcClient,
		logger: logger,
	}
}
