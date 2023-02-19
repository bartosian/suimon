package dashboardbuilder

import (
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"
)

type Cell struct {
	Config []container.Option
	Widget widgetapi.Widget
}

func NewCell(title string, widget widgetapi.Widget) *Cell {
	return &Cell{
		Config: []container.Option{
			container.FocusedColor(cell.ColorGreen),
			container.Border(linestyle.Light),
			container.BorderTitle(title),
			container.PlaceWidget(widget),
			container.AlignVertical(align.VerticalMiddle),
			container.AlignHorizontal(align.HorizontalCenter),
			container.BorderColor(cell.ColorRed),
			container.TitleColor(cell.ColorGreen),
		},
		Widget: widget,
	}
}

func NewTextCell(title string) *Cell {
	textWidget, err := newTextWidget()
	if err != nil {
		panic(err)
	}

	return NewCell(title, textWidget)
}

func NewProgressCell(title string) *Cell {
	gaugeWidget, err := newProgressWidget()
	if err != nil {
		panic(err)
	}

	return NewCell(title, gaugeWidget)
}

func NewDisplayCell(title string) *Cell {
	displayWidget, err := newDisplayWidget()
	if err != nil {
		panic(err)
	}

	return NewCell(title, displayWidget)
}

func newTextWidget() (*text.Text, error) {
	return text.New(text.RollContent(), text.WrapAtWords())
}

func newProgressWidget() (*gauge.Gauge, error) {
	return gauge.New(
		gauge.Height(3),
		gauge.Border(linestyle.Light, cell.FgColor(cell.ColorGreen)),
		gauge.Color(cell.ColorGreen),
		gauge.FilledTextColor(cell.ColorBlack),
		gauge.EmptyTextColor(cell.ColorWhite),
		gauge.HorizontalTextAlign(align.HorizontalCenter),
	)
}

func newDisplayWidget() (*segmentdisplay.SegmentDisplay, error) {
	return segmentdisplay.New()
}
