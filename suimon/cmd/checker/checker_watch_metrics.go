package checker

import (
	"sync"
	"time"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

const (
	watchHostsTimeout = 1 * time.Second
)

// WatchHosts begins monitoring the "Host" objects in the "Checker" struct instance passed as a pointer
// receiver, continuously checking their status and updating the dashboard and tables accordingly. This
// method runs indefinitely until the program is terminated, and does not return anything.
// Parameters: None.
// Returns: None.
func (checker *Checker) WatchHosts() {
	var (
		monitorsConfig = checker.suimonConfig.MonitorsConfig
		comparatorRPC  = checker.rpc[0]
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

						hosts[idx].SetPctProgress(enums.MetricTypeTxSyncProgress, comparatorRPC)
						hosts[idx].SetPctProgress(enums.MetricTypeCheckSyncProgress, comparatorRPC)
						hosts[idx].SetStatus(comparatorRPC)

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
