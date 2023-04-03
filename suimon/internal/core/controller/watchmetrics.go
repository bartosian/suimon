package controller

import (
	"sync"
	"time"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

const (
	watchHostsTimeout = 1 * time.Second
)

// Watch periodically retrieves the latest data from all active hosts and updates the CheckerController's internal state with the new data.
// The function retrieves data for each table type in parallel and displays a progress bar indicating the progress of the data retrieval process.
// Returns an error if the data cannot be retrieved from any of the active hosts or if there is an issue updating the CheckerController's internal state.
func (checker CheckerController) Watch() error {
	var (
		monitorsConfig = checker.suimonConfig.MonitorsConfig
		comparatorRPC  = checker.hosts.rpc[0]
		ticker         = time.NewTicker(watchHostsTimeout)
		ctx            = checker.DashboardBuilder.Ctx
		wg             sync.WaitGroup
	)

	defer ticker.Stop()

	var watch = func(hosts []host.Host, doneCH chan<- struct{}) {
		defer wg.Done()

		for {
			select {
			case <-ticker.C:
				for idx := range hosts {
					go func(idx int) {
						if err := hosts[idx].GetData(); err != nil {
							return
						}

						if err := hosts[idx].SetPctProgress(enums.MetricTypeTxSyncPercentage, comparatorRPC); err != nil {
							return
						}
						if err := hosts[idx].SetPctProgress(enums.MetricTypeCheckSyncPercentage, comparatorRPC); err != nil {
							return
						}

						hosts[idx].SetStatus(comparatorRPC)

						defer func() { doneCH <- struct{}{} }()
					}(idx)
				}

			case <-ctx.Done():
				return
			}
		}
	}

	if monitorsConfig.RPCTable.Display && len(checker.hosts.rpc) > 0 {
		doneCH := make(chan struct{}, len(checker.hosts.rpc))

		wg.Add(1)

		go watch(checker.hosts.rpc, doneCH)
	}

	if monitorsConfig.NodeTable.Display && len(checker.hosts.node) > 0 {
		doneCH := make(chan struct{}, len(checker.hosts.node))

		wg.Add(1)

		go watch(checker.hosts.node, doneCH)
	}

	if monitorsConfig.PeersTable.Display && len(checker.hosts.peers) > 0 {
		doneCH := make(chan struct{}, len(checker.hosts.peers))

		wg.Add(1)

		go watch(checker.hosts.peers, doneCH)
	}

	wg.Wait()

	return nil
}
