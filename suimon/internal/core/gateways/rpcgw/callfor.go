package rpcgw

import (
	"fmt"
	"time"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

func (gateway *Gateway) CallFor(method enums.RPCMethod) (result any, err error) {
	timeout := time.After(rpcClientTimeout)

	respChan := make(chan any)
	errChan := make(chan error)

	go func() {
		var response any

		if err := gateway.client.CallFor(gateway.ctx, &response, method.String()); err != nil {
			errChan <- err
		}

		respChan <- response
	}()

	select {
	case response := <-respChan:
		switch response.(type) {
		case nil:
			return nil, fmt.Errorf("failed to get response from RPC client")
		default:
			return response, nil
		}
	case err := <-errChan:
		return nil, fmt.Errorf("failed to get response from RPC client: %w", err)
	case <-timeout:
		return nil, fmt.Errorf("RPC call timed out after %s", rpcClientTimeout.String())
	}
}
