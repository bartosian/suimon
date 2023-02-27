package dashboards

import (
	"strconv"
	"strings"
	"time"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
)

const (
	dashboardName     = "💧 SUIMON: PRESS Q or ESC TO QUIT"
	logsWidgetMessage = "Please note that logs for sui-node can only be automatically captured from systemd services or Docker containers. If sui-node is run through other methods, such as manual start-up, logs may not be automatically captured and will need to be checked manually."
)

var (
	DashboardConfigSUI = []container.Option{
		container.Border(linestyle.Light),
		container.BorderColor(cell.ColorGreen),
		container.FocusedColor(cell.ColorGreen),
		container.AlignHorizontal(align.HorizontalCenter),
		container.AlignVertical(align.VerticalMiddle),
		container.BorderTitle(dashboardName),
		container.Focused(),
	}

	Rows = []grid.Element{
		0: NewRowPct(12,
			Columns[CellNameNodeStatus],
			Columns[CellNameNetworkStatus],
			Columns[CellNameCurrentEpoch],
			Columns[CellNameEpochEnd],
			Columns[CellNameUptime],
			Columns[CellNameConnectedPeers],
			Columns[CellNameVersion],
			Columns[CellNameCommit],
		),
		1: NewRowPct(12,
			Columns[CellNameTotalTransactions],
			Columns[CellNameLatestCheckpoint],
			Columns[CellNameHighestCheckpoint],
		),
		2: NewRowPct(12,
			Columns[CellNameTransactionsPerSecond],
			Columns[CellNameCheckpointsPerSecond],
			Columns[CellNameDatabaseSize],
			NewColumnPct(50,
				Columns[CellNameTXSyncProgress],
				Columns[CellNameCheckSyncProgress],
			),
		),
		3: NewRowPct(50,
			NewColumnPct(40,
				NewRowPct(42,
					Columns[CellNameEpochProgress],
					Columns[CellNameDiskUsage],
				),
				NewRowPct(42,
					Columns[CellNameCpuUsage],
					Columns[CellNameMemoryUsage],
				),
				NewRowPct(16,
					Columns[CellNameBytesSent],
					Columns[CellNameBytesReceived],
				),
			),
			Columns[CellNameNodeLogs],
		),
	}

	Columns = []grid.Element{
		CellNameNodeStatus:            NewColumnPct(5, Cells[CellNameNodeStatus].GetGridWidget()),
		CellNameNetworkStatus:         NewColumnPct(5, Cells[CellNameNetworkStatus].GetGridWidget()),
		CellNameTransactionsPerSecond: NewColumnPct(15, Cells[CellNameTransactionsPerSecond].GetGridWidget()),
		CellNameCheckpointsPerSecond:  NewColumnPct(15, Cells[CellNameCheckpointsPerSecond].GetGridWidget()),
		CellNameTotalTransactions:     NewColumnPct(33, Cells[CellNameTotalTransactions].GetGridWidget()),
		CellNameLatestCheckpoint:      NewColumnPct(33, Cells[CellNameLatestCheckpoint].GetGridWidget()),
		CellNameHighestCheckpoint:     NewColumnPct(33, Cells[CellNameHighestCheckpoint].GetGridWidget()),
		CellNameConnectedPeers:        NewColumnPct(10, Cells[CellNameConnectedPeers].GetGridWidget()),
		CellNameTXSyncProgress:        NewColumnPct(50, Cells[CellNameTXSyncProgress].GetGridWidget()),
		CellNameCheckSyncProgress:     NewColumnPct(50, Cells[CellNameCheckSyncProgress].GetGridWidget()),
		CellNameUptime:                NewColumnPct(10, Cells[CellNameUptime].GetGridWidget()),
		CellNameVersion:               NewColumnPct(20, Cells[CellNameVersion].GetGridWidget()),
		CellNameCommit:                NewColumnPct(10, Cells[CellNameCommit].GetGridWidget()),
		CellNameCurrentEpoch:          NewColumnPct(10, Cells[CellNameCurrentEpoch].GetGridWidget()),
		CellNameEpochProgress:         NewColumnPct(50, Cells[CellNameEpochProgress].GetGridWidget()),
		CellNameEpochEnd:              NewColumnPct(20, Cells[CellNameEpochEnd].GetGridWidget()),
		CellNameDiskUsage:             NewColumnPct(50, Cells[CellNameDiskUsage].GetGridWidget()),
		CellNameDatabaseSize:          NewColumnPct(20, Cells[CellNameDatabaseSize].GetGridWidget()),
		CellNameBytesSent:             NewColumnPct(50, Cells[CellNameBytesSent].GetGridWidget()),
		CellNameBytesReceived:         NewColumnPct(50, Cells[CellNameBytesReceived].GetGridWidget()),
		CellNameMemoryUsage:           NewColumnPct(50, Cells[CellNameMemoryUsage].GetGridWidget()),
		CellNameCpuUsage:              NewColumnPct(50, Cells[CellNameCpuUsage].GetGridWidget()),
		CellNameNodeLogs:              NewColumnPct(50, Cells[CellNameNodeLogs].GetGridWidget()),
	}

	Cells = []*Cell{
		CellNameNodeStatus:            NewCell("NODE", CellNameNodeStatus),
		CellNameNetworkStatus:         NewCell("NETWORK", CellNameNetworkStatus),
		CellNameTransactionsPerSecond: NewCell("TRANSACTIONS PER SECOND", CellNameTransactionsPerSecond),
		CellNameCheckpointsPerSecond:  NewCell("CHECKPOINTS PER SECOND", CellNameCheckpointsPerSecond),
		CellNameTotalTransactions:     NewCell("TOTAL TRANSACTIONS", CellNameTotalTransactions),
		CellNameLatestCheckpoint:      NewCell("LATEST CHECKPOINT", CellNameLatestCheckpoint),
		CellNameHighestCheckpoint:     NewCell("HIGHEST SYNCED CHECKPOINT", CellNameHighestCheckpoint),
		CellNameConnectedPeers:        NewCell("CONNECTED PEERS", CellNameConnectedPeers),
		CellNameTXSyncProgress:        NewCell("SYNC TRANSACTIONS STATUS", CellNameTXSyncProgress),
		CellNameCheckSyncProgress:     NewCell("SYNC CHECKPOINTS STATUS", CellNameCheckSyncProgress),
		CellNameUptime:                NewCell("UPTIME", CellNameUptime),
		CellNameVersion:               NewCell("VERSION", CellNameVersion),
		CellNameCommit:                NewCell("COMMIT", CellNameCommit),
		CellNameCurrentEpoch:          NewCell("CURRENT EPOCH", CellNameCurrentEpoch),
		CellNameEpochProgress:         NewCell("EPOCH PROGRESS", CellNameEpochProgress),
		CellNameEpochEnd:              NewCell("TIME TILL THE END OF EPOCH", CellNameEpochEnd),
		CellNameDiskUsage:             NewCell("DISK USAGE", CellNameDiskUsage),
		CellNameDatabaseSize:          NewCell("DATABASE SIZE", CellNameDatabaseSize),
		CellNameBytesSent:             NewCell("NETWORK BYTES SENT", CellNameBytesSent),
		CellNameBytesReceived:         NewCell("NETWORK BYTES RECEIVED", CellNameBytesReceived),
		CellNameMemoryUsage:           NewCell("MEMORY USAGE", CellNameMemoryUsage),
		CellNameCpuUsage:              NewCell("CPU USAGE", CellNameCpuUsage),
		CellNameNodeLogs:              NewCell("NODE LOGS", CellNameNodeLogs),
	}
)

type (
	Element interface {
		isElement()
	}

	Row struct {
		Height   int
		Elements []Element
	}

	Column struct {
		Width    int
		Elements []Element
	}

	Cell struct {
		Name    CellName
		Widget  widgetapi.Widget
		Options []container.Option
	}
)

func (Row) isElement() {}

func (Column) isElement() {}

func (Cell) isElement() {}

func NewCell(title string, name CellName) *Cell {
	return &Cell{
		Name:   name,
		Widget: newWidgetByCellName(name),
		Options: []container.Option{
			container.FocusedColor(cell.ColorGreen),
			container.Border(linestyle.Light),
			container.BorderTitle(title),
			container.BorderColor(cell.ColorGreen),
			container.AlignVertical(align.VerticalMiddle),
			container.AlignHorizontal(align.HorizontalCenter),
			container.TitleColor(cell.ColorGreen),
		},
	}
}

func NewRowPct(height int, elements ...grid.Element) grid.Element {
	return grid.RowHeightPerc(height, elements...)
}

func NewColumn(width int, elements ...grid.Element) grid.Element {
	return grid.ColWidthFixed(width, elements...)
}

func NewColumnPct(width int, elements ...grid.Element) grid.Element {
	return grid.ColWidthPerc(width, elements...)
}

type DonutWriteInput struct {
	Label      string
	Percentage int
	Color      cell.Color
}

func NewDonutInput(label string, pct int) DonutWriteInput {
	return DonutWriteInput{
		Label:      label,
		Percentage: pct,
	}
}

func (c *Cell) Write(value any, options any) {
	switch v := c.Widget.(type) {
	case *text.Text:
		valueString := value.(string)

		valueString = log.RemoveNonPrintableChars(valueString)
		if len(valueString) == 0 {
			return
		}

		if c.Name != CellNameNodeLogs {
			v.Reset()
		}

		textOptions := options.([]cell.Option)

		v.Write(valueString, text.WriteCellOpts(textOptions...))
	case *gauge.Gauge:
		valueInt := value.(int)

		gaugeOptions := options.([]gauge.Option)

		v.Percent(valueInt, gaugeOptions...)
	case *segmentdisplay.SegmentDisplay:
		var segments []*segmentdisplay.TextChunk

		segmentOptions := options.([]segmentdisplay.WriteOption)

		switch v := value.(type) {
		case int:
			chunk := strconv.Itoa(v)
			if chunk == "0" {
				chunk = dashboardLoadingBlinkValue()
			}

			segments = append(segments, segmentdisplay.NewChunk(chunk, segmentOptions...))
		case string:
			if v == "" {
				v = dashboardLoadingBlinkValue()
			}

			segments = append(segments, segmentdisplay.NewChunk(v, segmentOptions...))
		case []string:
			for idx, chunk := range v {
				if chunk == "" {
					chunk = dashboardLoadingBlinkValue()
				}

				segments = append(segments, segmentdisplay.NewChunk(chunk, segmentOptions[idx]))
			}
		}

		v.Write(segments)
	case *donut.Donut:
		valueInput := value.(DonutWriteInput)

		donutOptions := options.([]cell.Option)

		v.Percent(
			valueInput.Percentage,
			donut.Label(valueInput.Label, donutOptions...),
		)
	}
}

func (c *Cell) GetGridWidget() grid.Element {
	return grid.Widget(c.Widget, c.Options...)
}

func newProgressWidget() (*gauge.Gauge, error) {
	return gauge.New(
		gauge.Height(5),
		gauge.Border(linestyle.Light, cell.FgColor(cell.ColorGreen)),
		gauge.Color(cell.ColorGreen),
		gauge.FilledTextColor(cell.ColorBlack),
		gauge.EmptyTextColor(cell.ColorWhite),
		gauge.HorizontalTextAlign(align.HorizontalCenter),
		gauge.VerticalTextAlign(align.VerticalMiddle),
		gauge.Threshold(99, linestyle.Double, cell.FgColor(cell.ColorGreen), cell.Bold()),
	)
}

func newTextWidget() (*text.Text, error) {
	return text.New(text.RollContent(), text.WrapAtWords())
}

func newTextNoScrollWidget() (*text.Text, error) {
	return text.New(text.DisableScrolling(), text.WrapAtRunes())
}

func newDonutWidget(color cell.Color) (*donut.Donut, error) {
	return donut.New(
		donut.CellOpts(
			cell.FgColor(color),
			cell.Bold(),
		),
	)
}

func newButtonWidget(text string, color cell.Color, action func() error) (*button.Button, error) {
	return button.New(text, action,
		button.WidthFor("Submit"),
		button.FillColor(cell.ColorNumber(196)),
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

func dashboardLoadingValue() string {
	return strings.Repeat("-", 15)
}

func dashboardLoadingBlinkValue() string {
	inProgress := strings.Repeat("-", 15)
	second := time.Now().Second() % 10

	if second%2 == 0 {
		inProgress = strings.Repeat("\u0020", 15)
	}

	return inProgress
}

func newWidgetByCellName(name CellName) widgetapi.Widget {
	var err error

	switch name {
	case CellNameCheckSyncProgress, CellNameTXSyncProgress:
		var widget *gauge.Gauge

		if widget, err = newProgressWidget(); err == nil {
			widget.Percent(0)

			return widget
		}
	case CellNameNodeStatus, CellNameNetworkStatus:
		var widget *text.Text

		if widget, err = newTextNoScrollWidget(); err == nil {
			widget.Write(enums.StatusGrey.DashboardStatus(), text.WriteCellOpts(cell.FgColor(cell.ColorGray), cell.BgColor(cell.ColorGray)))

			return widget
		}
	case CellNameNodeLogs:
		var widget *text.Text

		if widget, err = newTextWidget(); err == nil {
			widget.Write(logsWidgetMessage, text.WriteCellOpts(cell.FgColor(cell.ColorWhite), cell.Bold()))

			return widget
		}
	case CellNameEpochProgress, CellNameDiskUsage, CellNameMemoryUsage, CellNameCpuUsage:
		var (
			widget *donut.Donut
			color  cell.Color
		)

		switch name {
		case CellNameEpochEnd:
			color = cell.ColorRGB24(51, 153, 102)
		case CellNameDiskUsage:
			color = cell.ColorRGB24(0, 255, 255)
		case CellNameCpuUsage:
			color = cell.ColorRGB24(255, 165, 0)
		case CellNameMemoryUsage:
			color = cell.ColorRGB24(255, 153, 204)
		}

		if widget, err = newDonutWidget(color); err == nil {
			widget.Percent(1, donut.Label("LOADING...", cell.FgColor(cell.ColorWhite), cell.Bold()))

			return widget
		}
	default:
		var widget *segmentdisplay.SegmentDisplay

		if widget, err = newDisplayWidget(); err == nil {
			widget.Write([]*segmentdisplay.TextChunk{
				segmentdisplay.NewChunk(dashboardLoadingValue(), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite))),
			})

			return widget
		}
	}

	panic(err)
}
