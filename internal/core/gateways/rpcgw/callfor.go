package rpcgw

import (
	"context"
	"fmt"

	"github.com/bartosian/suimon/internal/core/domain/enums"
)

type responseWithError struct {
	response any
	err      error
}

// CallFor executes an RPC method and returns the result or an error.
// The function sends an RPC request using the specified method and params, waits for the response, and handles timeouts.
func (gateway *Gateway) CallFor(method enums.RPCMethod, params ...interface{}) (result any, err error) {
	respChan := make(chan responseWithError)

	ctx, cancel := context.WithTimeout(gateway.ctx, rpcClientTimeout)
	defer cancel()

	go func() {
		var resp any

		err := gateway.client.CallFor(ctx, &resp, method.String(), params)

		if err != nil || resp == nil {
			respChan <- responseWithError{response: nil, err: fmt.Errorf("failed to get response from RPC client: %w", err)}
		} else {
			respChan <- responseWithError{response: resp, err: nil}
		}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("rpc call timed out: %w", ctx.Err())
	case result := <-respChan:
		if result.err != nil {
			return nil, result.err
		}

		return result.response, nil
	}
}
