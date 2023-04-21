package monitor

import (
	"context"
	"time"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

// WatchHostData monitors and updates the host's data periodically using the provided context.
// It watches the hosts based on the monitor's configuration.
func (c *Controller) WatchHostData(ctx context.Context, host domainhost.Host) error {
	rpcHost := c.hosts.rpc[0]

	ticker := time.NewTicker(dashboardRedrawInterval)
	defer ticker.Stop()

	errorChan := make(chan error)
	defer close(errorChan)

	var watch = func(host domainhost.Host) {
		for {
			select {
			case <-ticker.C:
				if err := host.GetMetrics(); err != nil {
					errorChan <- err

					return
				}

				if err := host.SetPctProgress(enums.MetricTypeTxSyncPercentage, rpcHost); err != nil {
					errorChan <- err

					return
				}
				if err := host.SetPctProgress(enums.MetricTypeCheckSyncPercentage, rpcHost); err != nil {
					errorChan <- err

					return
				}

				host.SetStatus(rpcHost)
			case <-ctx.Done():
				return
			}
		}
	}

	go watch(host)

	select {
	case err := <-errorChan:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
