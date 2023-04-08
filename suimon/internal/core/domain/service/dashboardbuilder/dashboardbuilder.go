package dashboardbuilder

import (
	"context"
	"os"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder/config"
)

type Builder struct {
	Ctx       context.Context
	Terminal  *termbox.Terminal
	Dashboard *container.Container
	Cells     []*config.Cell
	Quitter   func(k *terminalapi.Keyboard)
}

func NewBuilder() (*Builder, error) {
	terminal, err := termbox.New()
	if err != nil {
		return nil, err
	}

	dashboard, err := initDashboard(terminal)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Builder{
		Ctx:       ctx,
		Terminal:  terminal,
		Dashboard: dashboard,
		Cells:     config.Cells,
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
	builder := grid.New()
	rows := config.Rows

	builder.Add(rows...)

	options, err := builder.Build()
	if err != nil {
		return nil, err
	}
	dashboard := config.DashboardConfigNode
	dashboard = append(dashboard, options...)

	return container.New(
		terminal,
		dashboard...,
	)
}
