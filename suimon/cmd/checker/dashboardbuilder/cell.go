package dashboardbuilder

import (
	"strconv"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder/dashboards"
)

const dashboardInProgress = "LOAD"

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

type DonutWriteInput struct {
	Label      string
	Percentage int
}

func NewDonutInput(label string, pct int) DonutWriteInput {
	return DonutWriteInput{
		Label:      label,
		Percentage: pct,
	}
}

func (c *Cell) Write(value any) {
	switch v := c.Widget.(type) {
	case *text.Text:
		valueString := value.(string)

		if valueString == "" {
			valueString = dashboardInProgress
		}

		v.Reset()
		v.Write(valueString, text.WriteCellOpts(cell.Bold()))
	case *gauge.Gauge:
		valueInt := value.(int)

		v.Percent(valueInt)
	case *segmentdisplay.SegmentDisplay:
		var chunkValue string

		switch v := value.(type) {
		case int:
			chunkValue = strconv.Itoa(v)
		case string:
			chunkValue = v
		}

		if chunkValue == "" || chunkValue == "0" {
			chunkValue = dashboardInProgress
		}

		v.Write([]*segmentdisplay.TextChunk{
			segmentdisplay.NewChunk(chunkValue),
		})
	case *donut.Donut:
		valueInput := value.(DonutWriteInput)

		v.Percent(
			valueInput.Percentage,
			donut.Label(valueInput.Label, cell.FgColor(cell.ColorGreen), cell.Bold()),
		)
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

func NewDonutCell(title string) *Cell {
	donutWidget, err := newDonutWidget()
	if err != nil {
		panic(err)
	}

	return NewCell(title, donutWidget)
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
		gauge.VerticalTextAlign(align.VerticalMiddle),
	)
}

func newDisplayWidget() (*segmentdisplay.SegmentDisplay, error) {
	return segmentdisplay.New(
		segmentdisplay.MaximizeDisplayedText(),
		segmentdisplay.AlignHorizontal(align.HorizontalCenter),
		segmentdisplay.AlignVertical(align.VerticalMiddle),
		segmentdisplay.GapPercent(10),
	)
}

func newDonutWidget() (*donut.Donut, error) {
	return donut.New(
		donut.CellOpts(
			cell.FgColor(cell.ColorGreen),
			cell.Bold(),
		),
	)
}

func initCells() []*Cell {
	var cells = make([]*Cell, len(dashboards.ColumnConfigSUI))

	for name, config := range dashboards.ColumnConfigSUI {
		var (
			nameEnum   = dashboards.CellName(name)
			nameString = config.Name
			cell       *Cell
		)

		switch nameEnum {
		case dashboards.CellNameCheckSyncProgress, dashboards.CellNameTXSyncProgress:
			cell = NewProgressCell(nameString)
		case dashboards.CellNameStatus:
			cell = NewTextCell(nameString)
		case dashboards.CellNameEpoch:
			cell = NewDonutCell(nameString)
		default:
			cell = NewDisplayCell(nameString)
		}

		cells[name] = cell
	}

	return cells
}
