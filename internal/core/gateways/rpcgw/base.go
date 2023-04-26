package rpcgw

import (
	"context"
	"net/http"
	"time"

	"github.com/ybbus/jsonrpc/v3"

	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
	"github.com/bartosian/sui_helpers/suimon/internal/core/ports"
)

const rpcClientTimeout = 3 * time.Second

type Gateway struct {
	ctx        context.Context
	url        string
	client     jsonrpc.RPCClient
	cliGateway *cligw.Gateway
}

func NewGateway(cliGW *cligw.Gateway, url string) ports.RPCGateway {
	httpClient := &http.Client{
		Timeout: rpcClientTimeout,
	}

	opts := &jsonrpc.RPCClientOpts{
		HTTPClient: httpClient,
	}

	rpcClient := jsonrpc.NewClientWithOpts(url, opts)

	return &Gateway{
		ctx:        context.Background(),
		url:        url,
		client:     rpcClient,
		cliGateway: cliGW,
	}
}
