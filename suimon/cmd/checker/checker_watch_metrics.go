package checker

import (
	"context"
	"sync"
	"time"
)

const (
	watchHostsTimeout = 1 * time.Second
)

func (checker *Checker) WatchHosts(ctx context.Context) {
	var (
		monitorsConfig = checker.suimonConfig.MonitorsConfig
		ticker         = time.NewTicker(watchHostsTimeout)
		hosts          []Host
	)

	if monitorsConfig.RPCTable.Display && len(checker.rpc) > 0 {
		hosts = append(hosts, checker.rpc...)
	}

	if monitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		hosts = append(hosts, checker.node...)
	}

	if monitorsConfig.PeersTable.Display && len(checker.peers) > 0 {
		hosts = append(hosts, checker.peers...)
	}

	for {
		select {
		case <-ticker.C:
			var wg sync.WaitGroup

			for idx := range hosts {
				wg.Add(1)

				go func(idx int) {
					defer wg.Done()

					hosts[idx].GetData()
				}(idx)
			}

			wg.Wait()
		case <-ctx.Done():
			return
		}
	}
}

func (host *Host) GetData() {
	host.stateMutex.Lock()

	defer host.stateMutex.Unlock()

	doneCH := make(chan struct{})

	defer close(doneCH)

	go func() {
		host.GetTotalTransactionNumber()

		doneCH <- struct{}{}
	}()

	go func() {
		host.GetLatestCheckpoint()

		doneCH <- struct{}{}
	}()

	go func() {
		host.GetMetrics()

		doneCH <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		<-doneCH
	}
}
