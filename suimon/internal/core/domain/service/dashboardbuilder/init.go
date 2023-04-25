package dashboardbuilder

import (
	"fmt"
	"os"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder/dashboards"
)

// Init initializes the dashboard by fetching the cells, columns, and rows
// configurations from the `dashboards` package and using them to build a new
// `grid` with the `grid.New()` method. It then uses the built grid to create a
// new dashboard using the `container.New()` method. The dashboard instance is
// stored in the `db.dashboard` field for later use.
func (db *Builder) Init() (err error) {
	// Use a deferred function to call db.TearDown() if there were errors or panics
	defer func() {
		if err != nil {
			db.tearDown()
		}

		if err := recover(); err != nil {
			// Handle the panic by logging the error and exiting the program
			db.tearDown()

			db.cliGateway.Error(fmt.Sprintf("panic: %v", err))

			os.Exit(1)
		}
	}()

	cellsConfig, err := dashboards.GetCellsConfig(db.tableType)
	if err != nil {
		return err
	}

	cells, err := dashboards.GetCells(cellsConfig)
	if err != nil {
		return err
	}

	db.cells = cells

	columnsConfig, err := dashboards.GetColumnsConfig(db.tableType)
	if err != nil {
		return err
	}

	columns, err := dashboards.GetColumns(columnsConfig, cells)
	if err != nil {
		return err
	}

	rowsConfig, err := dashboards.GetRowsConfig(db.tableType)
	if err != nil {
		return err
	}

	rows, err := dashboards.GetRows(rowsConfig, cells, columns)
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
