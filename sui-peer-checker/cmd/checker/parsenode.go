package checker

import (
	"errors"
	"strings"
	"sync"
)

func (checker *Checker) parseNode() error {
	addressRPC, addressMetrics := checker.nodeYaml.JsonRPCAddress, checker.nodeYaml.MetricsAddress

	addressPort := strings.Split(addressRPC, addressSeparator)
	if len(addressPort) != 2 {
		return errors.New("invalid json-rpc-address in config file")
	}

	metricsAddressPort := strings.Split(addressMetrics, addressSeparator)

	if len(addressPort) != 2 {
		return errors.New("invalid metrics-address in config file")
	}

	node := newNode(
		checker.geoDbClient,
		checker.httpClient,
		addressPort[0],
		addressPort[1],
		metricsAddressPort[1],
	)

	if err := node.Parse(); err != nil {
		return err
	}

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()

		node.GetTotalTransactionNumber()
	}()

	go func() {
		defer wg.Done()

		node.GetMetrics()
	}()

	wg.Wait()

	checker.node = *node

	return nil
}
