package checker

import (
	"errors"
	"sync"
	"time"

	"github.com/hashicorp/go-multierror"
)

const (
	httpClientTimeout = 3 * time.Second
)

func (checker *Checker) ParseData() error {
	suimonConfig, nodeConfig := checker.suimonConfig, checker.nodeConfig

	if len(nodeConfig.P2PConfig.SeedPeers) == 0 {
		return errors.New("no peers found in config file")
	}

	var (
		wg         sync.WaitGroup
		errChan    = make(chan error)
		errCounter int
		err        error
	)

	monitorsConfig := suimonConfig.MonitorsConfig

	// parse data for the RPC table
	if monitorsConfig.RPCTable.Display {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := checker.parseRPCHosts(); err != nil {
				errChan <- err
			}
		}()
	}

	// parse data for the NODE table
	if monitorsConfig.NodeTable.Display {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := checker.parseNode(); err != nil {
				errChan <- err
			}
		}()
	}

	// parse data for the PEERS table
	if monitorsConfig.PeersTable.Display {
		wg.Add(1)

		go func() {
			defer wg.Done()

			if err := checker.parsePeers(); err != nil {
				errChan <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for parseErr := range errChan {
		err = multierror.Append(err, parseErr)

		errCounter++
	}

	if errCounter == 3 {
		return err
	}

	return nil
}
