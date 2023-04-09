package dashboardbuilder

import (
	"fmt"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder/config"
)

// Init initializes the Builder's dashboard by adding the rows to the grid builder
// and setting the dashboard configuration. If an error occurs during initialization,
// it returns an error.
func (db *Builder) Init() error {
	builder := grid.New()
	rows := config.Rows

	builder.Add(rows...)

	options, err := builder.Build()
	if err != nil {
		return err
	}

	dashboardConfig := append(config.DashboardConfigNode, options...)

	dashboard, err := container.New(db.terminal, dashboardConfig...)
	if err != nil {
		return fmt.Errorf("failed to initialize dashboard: %w", err)
	}

	db.dashboard = dashboard

	return nil
}
