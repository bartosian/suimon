package dashboardbuilder

import (
	"fmt"
	"time"

	"github.com/mum4k/termdash"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	domainhost "github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

const renderInterval = 1 * time.Second

func (db *Builder) Render(host domainhost.Host) error {
	ticker := time.NewTicker(renderInterval)
	defer ticker.Stop()

	var render = func(host domainhost.Host) {
		for {
			select {
			case <-ticker.C:
				for idx, cell := range db.cells {
					metricName := enums.CellName(idx)

					metric := host.Metrics.GetValue(cellName)
					options := checker.getOptionsForDashboardCell(cellName)

					cell.Write(metric, options)
				}
			case <-db.ctx.Done():
				return
			}
		}
	}

	go render(host)

	err := termdash.Run(db.ctx, db.terminal, db.dashboard, termdash.KeyboardSubscriber(db.quitter))
	if err != nil {
		return fmt.Errorf("failed to run terminal dashboard: %w", err)
	}

	select {
	case err := <-errorChan:
		return err
	case <-db.ctx.Done():
		return nil
	}
}
