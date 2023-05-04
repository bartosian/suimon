package dashboardbuilder

import (
	"context"
	"fmt"
	"os"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/keyboard"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder/dashboards"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
)

type Builder struct {
	ctx        context.Context
	tableType  enums.TableType
	cliGateway *cligw.Gateway
	terminal   *termbox.Terminal
	dashboard  *container.Container
	host       host.Host
	cells      dashboards.Cells
	quitter    func(k *terminalapi.Keyboard)
}

// NewBuilder creates a new Builder instance with the provided CLI gateway.
// It initializes the termbox terminal and dashboard, and sets up a context and quitter function.
// If an error occurs during initialization, it returns an error.
func NewBuilder(tableType enums.TableType, host host.Host, cliGateway *cligw.Gateway) (*Builder, error) {
	terminal, err := termbox.New()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize termbox terminal: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Builder{
		ctx:        ctx,
		tableType:  tableType,
		cliGateway: cliGateway,
		terminal:   terminal,
		host:       host,
		quitter: func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' || k.Key == keyboard.KeyEsc || k.Key == keyboard.KeyCtrlC {
				terminal.Close()
				cancel()

				os.Exit(0)
			}
		},
	}, nil
}

// The tearDown function closes the Builder's terminal and cancels its context.
func (db *Builder) tearDown() {
	db.ctx.Done()
	db.terminal.Close()
}
