package rpcgw

import (
	"context"
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
	rpcClient := jsonrpc.NewClient(url)

	return &Gateway{
		url:    url,
		ctx:    context.Background(),
		client: rpcClient,
		logger: logger,
	}
}
