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

func initDashboard(terminal *termbox.Terminal) (*container.Container, error) {
	var (
		builder   = grid.New()
		dashboard = dashboards.DashboardConfigSUI
		rows      = dashboards.Rows
	)

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
