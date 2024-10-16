package dashboardbuilder

import (
	"fmt"
	"os"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"

	"github.com/bartosian/suimon/internal/core/domain/service/dashboardbuilder/dashboards"
)

// Init initializes the dashboard by fetching and configuring the grid.
func (db *Builder) Init() (err error) {
	defer func() {
		if err != nil {
			db.tearDown()
		}

		if recoverErr := recover(); recoverErr != nil {
			db.tearDown()
			db.cliGateway.Error(fmt.Sprintf("panic: %v", recoverErr))
			os.Exit(1)
		}
	}()

	if setupErr := db.setupDashboard(); setupErr != nil {
		return setupErr
	}

	return nil
}

// setupDashboard is the main method that orchestrates the dashboard setup.
func (db *Builder) setupDashboard() error {
	cells, err := db.loadCells()
	if err != nil {
		return err
	}

	db.cells = cells

	columns, err := db.loadColumns(cells)
	if err != nil {
		return err
	}

	rows, err := db.loadRows(columns)
	if err != nil {
		return err
	}

	builder := grid.New()
	builder.Add(rows...)

	options, err := builder.Build()
	if err != nil {
		return err
	}

	return db.createDashboard(options)
}

// loadCells fetches the cells configuration and builds the cells.
func (db *Builder) loadCells() (dashboards.Cells, error) {
	cellsConfig, err := dashboards.GetCellsConfig(db.tableType)
	if err != nil {
		return nil, err
	}

	cells, err := dashboards.GetCells(cellsConfig)
	if err != nil {
		return nil, err
	}

	return cells, nil
}

// loadColumns fetches the columns configuration and builds the columns based on the cells.
func (db *Builder) loadColumns(cells dashboards.Cells) (dashboards.Columns, error) {
	columnsConfig, err := dashboards.GetColumnsConfig(db.tableType)
	if err != nil {
		return nil, err
	}

	columns, err := dashboards.GetColumns(columnsConfig, cells)
	if err != nil {
		return nil, err
	}

	return columns, nil
}

// loadRows fetches the rows configuration and builds the rows based on the columns.
func (db *Builder) loadRows(columns dashboards.Columns) ([]grid.Element, error) {
	rowsConfig, err := dashboards.GetRowsConfig(db.tableType)
	if err != nil {
		return nil, err
	}

	rows, err := dashboards.GetRows(rowsConfig, columns)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// createDashboard creates the dashboard using the built grid and terminal.
func (db *Builder) createDashboard(options []container.Option) error {
	dashboardConfig := append([]container.Option{}, dashboards.DashboardConfigDefault...)
	dashboardConfig = append(dashboardConfig, options...)

	dashboard, err := container.New(db.terminal, dashboardConfig...)
	if err != nil {
		return fmt.Errorf("failed to initialize dashboard: %w", err)
	}

	db.dashboard = dashboard

	return nil
}
