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

// CallFor makes an RPC call for the specified method with the given parameters.
// It returns the result of the RPC call and an error if any.
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
