package dashboardbuilder

import (
	"context"
	"os"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
)

type DashboardBuilder struct {
	Ctx       context.Context
	Terminal  *termbox.Terminal
	Dashboard *container.Container
	Cells     []*Cell
	Quitter   func(k *terminalapi.Keyboard)
}

func NewDashboardBuilder() (*DashboardBuilder, error) {
	var (
		terminal  *termbox.Terminal
		dashboard *container.Container
		cells     []*Cell
		err       error
	)

	ctx, cancel := context.WithCancel(context.Background())

	if terminal, err = termbox.New(); err != nil {
		return nil, err
	}

	cells = initCells()

	if dashboard, err = initDashboard(terminal, cells); err != nil {
		return nil, err
	}

	return &DashboardBuilder{
		Ctx:       ctx,
		Terminal:  terminal,
		Dashboard: dashboard,
		Cells:     cells,
		Quitter: func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc {
				terminal.Close()
				cancel()
				os.Exit(0)
			}
		},
	}, nil
}

func initDashboard(terminal *termbox.Terminal, cells []*Cell) (*container.Container, error) {
	var (
		builder   = grid.New()
		columns   = make([]grid.Element, len(dashboards.ColumnConfigSUI))
		rows      = make([]grid.Element, len(dashboards.RowConfigSUI))
		dashboard = dashboards.DashboardConfigSUI
	)

	for idx, columnConfig := range dashboards.ColumnConfigSUI {
		var (
			width      = columnConfig.Width
			cell       = cells[idx]
			widget     = cell.Widget
			cellConfig = cell.Config
		)

		columns[idx] = grid.ColWidthFixed(width, grid.Widget(widget, cellConfig...))
	}

	for idx, rowConfig := range dashboards.RowConfigSUI {
		var (
			height     = rowConfig.Height
			colNames   = rowConfig.Columns
			rowColumns = make([]grid.Element, len(colNames))
		)

		for idx, columnName := range colNames {
			rowColumns[idx] = columns[columnName]
		}

		rowColumns = append(rowColumns, grid.ColWidthFixed(0))

		rows[idx] = grid.RowHeightFixed(height, rowColumns...)
	}

	rows = append(rows, grid.RowHeightFixed(0))

	builder.Add(rows...)

	config, err := builder.Build()
	if err != nil {
		return nil, err
	}

	dashboard = append(dashboard, config...)

	return container.New(
		terminal,
		dashboard...,
	)
}
