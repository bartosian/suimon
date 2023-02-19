package checker

import (
	"sync"
	"time"
)

const (
	watchHostsTimeout = 1 * time.Second
)

func (checker *Checker) WatchHosts() {
	var (
		monitorsConfig = checker.suimonConfig.MonitorsConfig
		ticker         = time.NewTicker(watchHostsTimeout)
		ctx            = checker.DashboardBuilder.Ctx
		wg             sync.WaitGroup
	)

	defer ticker.Stop()

	var watch = func(hosts []Host) {
		defer wg.Done()

		doneCH := make(chan struct{}, len(hosts))

		for {
			select {
			case <-ticker.C:
				for idx := range hosts {
					go func(idx int) {
						hosts[idx].GetData()

						doneCH <- struct{}{}
					}(idx)
				}

				for i := 0; i < len(hosts); i++ {
					<-doneCH
				}
			case <-ctx.Done():
				return
			}
		}
	}

	if monitorsConfig.RPCTable.Display && len(checker.rpc) > 0 {
		wg.Add(1)

		go watch(checker.rpc)
	}

	if monitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		wg.Add(1)

		go watch(checker.node)
	}

	if monitorsConfig.PeersTable.Display && len(checker.peers) > 0 {
		wg.Add(1)

		go watch(checker.peers)
	}

	wg.Wait()
}
