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
	Cells     []*dashboards.Cell
	Quitter   func(k *terminalapi.Keyboard)
}

// NewDashboardBuilder returns a new instance of a DashboardBuilder, which can be used to construct suimon dashboards.
// Returns: a pointer to the new DashboardBuilder instance, and any error encountered during the construction process.
func NewDashboardBuilder() (*DashboardBuilder, error) {
	var (
		terminal  *termbox.Terminal
		dashboard *container.Container
		err       error
	)

	ctx, cancel := context.WithCancel(context.Background())

	if terminal, err = termbox.New(); err != nil {
		return nil, err
	}

	if dashboard, err = initDashboard(terminal); err != nil {
		return nil, err
	}

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

// initDashboard initializes the dashboard UI by setting up the container and initializing the top-level views.
// Parameters: a `terminal` object as a parameter, which is used to display the UI in the terminal.
// Returns: a pointer to a `container.Container` object and an error if there was an issue initializing the dashboard.
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
