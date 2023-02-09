package checker

import (
	"errors"
	externalip "github.com/glendc/go-external-ip"
	"net"
	"strings"
	"sync"
)

func (checker *Checker) parseNode() error {
	addressRPC, addressMetrics := checker.nodeYaml.JsonRPCAddress, checker.nodeYaml.MetricsAddress

	addressRPCInfo := strings.Split(addressRPC, addressSeparator)
	if len(addressRPCInfo) != 2 {
		return errors.New("invalid json-rpc-address in config file")
	}

	addressMetricsInfo := strings.Split(addressMetrics, addressSeparator)
	if len(addressMetricsInfo) != 2 {
		return errors.New("invalid metrics-address in config file")
	}

	publicIP := getPublicIP()

	node := newNode(
		checker.geoDbClient,
		checker.httpClient,
		publicIP.String(),
		addressRPCInfo[1],
		addressMetricsInfo[1],
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

func getPublicIP() net.IP {
	consensus := externalip.DefaultConsensus(nil, nil)

	ip, err := consensus.ExternalIP()
	if err != nil {
		return nil
	}

	return ip
}
