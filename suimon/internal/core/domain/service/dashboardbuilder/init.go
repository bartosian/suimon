package dashboardbuilder

import (
	"errors"
	"fmt"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder/dashboards"
)

// Init initializes the dashboard by fetching the cells, columns, and rows
// configurations from the `dashboards` package and using them to build a new
// `grid` with the `grid.New()` method. It then uses the built grid to create a
// new dashboard using the `container.New()` method. The dashboard instance is
// stored in the `db.dashboard` field for later use.
func (db *Builder) Init() error {
	hosts := db.hosts

	if len(hosts) == 0 {
		return errors.New("hosts are not initialized")
	}

	cellsConfig := dashboards.GetCellsConfig(db.tableType)
	cells, err := dashboards.GetCells(cellsConfig)
	if err != nil {
		return err
	}

	db.cells = cells

	columnsConfig := dashboards.GetColumnsConfig(db.tableType)
	columns, err := dashboards.GetColumns(columnsConfig, cells)
	if err != nil {
		return err
	}

	rowsConfig := dashboards.GetRowsConfig(db.tableType)
	rows, err := dashboards.GetRows(rowsConfig, columns)
	if err != nil {
		return err
	}

	builder := grid.New()
	builder.Add(rows...)

	options, err := builder.Build()
	if err != nil {
		return err
	}

	dashboardConfig := append(dashboards.DashboardConfigDefault, options...)

	dashboard, err := container.New(db.terminal, dashboardConfig...)
	if err != nil {
		return fmt.Errorf("failed to initialize dashboard: %w", err)
	}

	db.dashboard = dashboard

	return nil
}
