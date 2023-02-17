package dashboardbuilder

import (
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/terminal/termbox"
)

type DashboardBuilder struct {
	terminal  *termbox.Terminal
	dashboard *container.Container
}
