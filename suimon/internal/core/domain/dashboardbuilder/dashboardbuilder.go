package dashboardbuilder

import (
	"context"
	"os"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/dashboardbuilder/dashboards"
)

type DashboardBuilder struct {
	Ctx       context.Context
	Terminal  *termbox.Terminal
	Dashboard *container.Container
	Cells     []*dashboards.Cell
	Quitter   func(k *terminalapi.Keyboard)
}

// NewDashboardBuilder creates a new DashboardBuilder instance that can be used to generate and display dashboards.
// The function initializes a terminal instance and a dashboard container using the termbox library, and initializes the DashboardBuilder's internal state with the configured cells.
// Returns a pointer to the DashboardBuilder instance and an error if there is an issue creating the terminal or initializing the dashboard.
func NewDashboardBuilder() (*DashboardBuilder, error) {
	var (
		terminal  *termbox.Terminal
		dashboard *container.Container
		err       error
	)

	if terminal, err = termbox.New(); err != nil {
		return nil, err
	}

	if dashboard, err = initDashboard(terminal); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &DashboardBuilder{
		Ctx:       ctx,
		Terminal:  terminal,
		Dashboard: dashboard,
		Cells:     dashboards.Cells,
		Quitter: func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc {
				terminal.Close()
				cancel()
				os.Exit(0)
			}
		},
	}, nil
}

// initDashboard initializes a new dashboard container using the given terminal instance and the configured rows and dashboard configuration.
// The function creates a new grid builder and adds the configured rows to the builder.
// It then builds the grid configuration and adds it to the dashboard configuration.
// Finally, the function creates a new container instance using the terminal and dashboard configuration.
// Returns a pointer to the container instance and an error if there is an issue initializing the dashboard.
func initDashboard(terminal *termbox.Terminal) (*container.Container, error) {
	var (
		builder   = grid.New()
		dashboard = dashboards.DashboardConfigNode
		rows      = dashboards.Rows
		config    []container.Option
		err       error
	)

	builder.Add(rows...)

	if config, err = builder.Build(); err != nil {
		return nil, err
	}

	dashboard = append(dashboard, config...)

	return container.New(
		terminal,
		dashboard...,
	)
}
