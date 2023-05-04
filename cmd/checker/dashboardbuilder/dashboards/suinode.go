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
	"github.com/mum4k/termdash/widgets/donut"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/suimon/cmd/checker/enums"
	"github.com/bartosian/suimon/pkg/log"
)

const (
	dashboardName     = "💧 SUIMON: PRESS Q or ESC TO QUIT"
	logsWidgetMessage = "Looking for the sui-node process to stream the logs.\n\n"
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
			Columns[CellNameConnectedPeers],
			Columns[CellNameVersion],
			Columns[CellNameCommit],
		),
		1: NewRowPct(12,
			Columns[CellNameUptime],
			Columns[CellNameTotalTransactions],
			Columns[CellNameLatestCheckpoint],
			Columns[CellNameHighestCheckpoint],
		),
		2: NewRowPct(12,
			Columns[CellNameBytesSent],
			Columns[CellNameBytesReceived],
			Columns[CellNameTransactionsPerSecond],
			Columns[CellNameCheckpointsPerSecond],
			Columns[CellNameDatabaseSize],
		),
		3: NewRowPct(50,
			NewColumnPct(40,
				NewRowPct(50,
					Columns[CellNameEpochProgress],
					Columns[CellNameDiskUsage],
				),
				NewRowPct(50,
					Columns[CellNameCpuUsage],
					Columns[CellNameMemoryUsage],
				),
			),
			NewColumnPct(60,
				NewRowPct(15,
					Columns[CellNameTXSyncProgress],
					Columns[CellNameCheckSyncProgress],
				),
				NewRowPct(20, Columns[CellNameTPSTracker]),
				NewRowPct(20, Columns[CellNameCPSTracker]),
				NewRowPct(45, Columns[CellNameNodeLogs]),
			),
		),
	}

	Columns = []grid.Element{
		CellNameNodeStatus:            NewColumnPct(5, Cells[CellNameNodeStatus].GetGridWidget()),
		CellNameNetworkStatus:         NewColumnPct(5, Cells[CellNameNetworkStatus].GetGridWidget()),
		CellNameTransactionsPerSecond: NewColumnPct(16, Cells[CellNameTransactionsPerSecond].GetGridWidget()),
		CellNameCheckpointsPerSecond:  NewColumnPct(15, Cells[CellNameCheckpointsPerSecond].GetGridWidget()),
		CellNameTotalTransactions:     NewColumnPct(28, Cells[CellNameTotalTransactions].GetGridWidget()),
		CellNameLatestCheckpoint:      NewColumnPct(28, Cells[CellNameLatestCheckpoint].GetGridWidget()),
		CellNameHighestCheckpoint:     NewColumnPct(28, Cells[CellNameHighestCheckpoint].GetGridWidget()),
		CellNameConnectedPeers:        NewColumnPct(10, Cells[CellNameConnectedPeers].GetGridWidget()),
		CellNameTXSyncProgress:        NewColumnPct(50, Cells[CellNameTXSyncProgress].GetGridWidget()),
		CellNameCheckSyncProgress:     NewColumnPct(50, Cells[CellNameCheckSyncProgress].GetGridWidget()),
		CellNameUptime:                NewColumnPct(16, Cells[CellNameUptime].GetGridWidget()),
		CellNameVersion:               NewColumnPct(21, Cells[CellNameVersion].GetGridWidget()),
		CellNameCommit:                NewColumnPct(10, Cells[CellNameCommit].GetGridWidget()),
		CellNameCurrentEpoch:          NewColumnPct(10, Cells[CellNameCurrentEpoch].GetGridWidget()),
		CellNameEpochProgress:         NewColumnPct(50, Cells[CellNameEpochProgress].GetGridWidget()),
		CellNameEpochEnd:              NewColumnPct(20, Cells[CellNameEpochEnd].GetGridWidget()),
		CellNameDiskUsage:             NewColumnPct(50, Cells[CellNameDiskUsage].GetGridWidget()),
		CellNameDatabaseSize:          NewColumnPct(20, Cells[CellNameDatabaseSize].GetGridWidget()),
		CellNameBytesSent:             NewColumnPct(20, Cells[CellNameBytesSent].GetGridWidget()),
		CellNameBytesReceived:         NewColumnPct(20, Cells[CellNameBytesReceived].GetGridWidget()),
		CellNameMemoryUsage:           NewColumnPct(50, Cells[CellNameMemoryUsage].GetGridWidget()),
		CellNameCpuUsage:              NewColumnPct(50, Cells[CellNameCpuUsage].GetGridWidget()),
		CellNameNodeLogs:              NewColumnPct(50, Cells[CellNameNodeLogs].GetGridWidget()),
		CellNameButtonQuit:            NewColumnPct(25, Cells[CellNameButtonQuit].GetGridWidget()),
		CellNameTPSTracker:            NewColumnPct(99, Cells[CellNameTPSTracker].GetGridWidget()),
		CellNameCPSTracker:            NewColumnPct(99, Cells[CellNameCPSTracker].GetGridWidget()),
	}

	Cells = []*Cell{
		CellNameNodeStatus:            NewCell(CellNameNodeStatus),
		CellNameNetworkStatus:         NewCell(CellNameNetworkStatus),
		CellNameTransactionsPerSecond: NewCell(CellNameTransactionsPerSecond),
		CellNameCheckpointsPerSecond:  NewCell(CellNameCheckpointsPerSecond),
		CellNameTotalTransactions:     NewCell(CellNameTotalTransactions),
		CellNameLatestCheckpoint:      NewCell(CellNameLatestCheckpoint),
		CellNameHighestCheckpoint:     NewCell(CellNameHighestCheckpoint),
		CellNameConnectedPeers:        NewCell(CellNameConnectedPeers),
		CellNameTXSyncProgress:        NewCell(CellNameTXSyncProgress),
		CellNameCheckSyncProgress:     NewCell(CellNameCheckSyncProgress),
		CellNameUptime:                NewCell(CellNameUptime),
		CellNameVersion:               NewCell(CellNameVersion),
		CellNameCommit:                NewCell(CellNameCommit),
		CellNameCurrentEpoch:          NewCell(CellNameCurrentEpoch),
		CellNameEpochProgress:         NewCell(CellNameEpochProgress),
		CellNameEpochEnd:              NewCell(CellNameEpochEnd),
		CellNameDiskUsage:             NewCell(CellNameDiskUsage),
		CellNameDatabaseSize:          NewCell(CellNameDatabaseSize),
		CellNameBytesSent:             NewCell(CellNameBytesSent),
		CellNameBytesReceived:         NewCell(CellNameBytesReceived),
		CellNameMemoryUsage:           NewCell(CellNameMemoryUsage),
		CellNameCpuUsage:              NewCell(CellNameCpuUsage),
		CellNameNodeLogs:              NewCell(CellNameNodeLogs),
		CellNameButtonQuit:            NewCell(CellNameButtonQuit),
		CellNameTPSTracker:            NewCell(CellNameTPSTracker),
		CellNameCPSTracker:            NewCell(CellNameCPSTracker),
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

func NewCell(name CellName) *Cell {
	return &Cell{
		Name:   name,
		Widget: newWidgetByCellName(name),
		Options: []container.Option{
			container.FocusedColor(cell.ColorGreen),
			container.Border(linestyle.Light),
			container.BorderTitle(name.CellNameString()),
			container.BorderColor(cell.ColorRed),
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
		textOptions := options.([]cell.Option)

		if valueString = log.RemoveNonPrintableChars(valueString); len(valueString) == 0 {
			return
		}

		if c.Name != CellNameNodeLogs {
			v.Reset()
		}

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

			segments = append(segments, segmentdisplay.NewChunk(chunk, segmentOptions...))
		case string:
			if v == "" {
				v = DashboardLoadingBlinkValue()
			}

			segments = append(segments, segmentdisplay.NewChunk(v, segmentOptions...))
		case []string:
			for idx, chunk := range v {
				if chunk == "" {
					chunk = DashboardLoadingBlinkValue()
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
	case *sparkline.SparkLine:
		second := time.Now().Second()
		if second%5 != 0 {
			return
		}

		valueInput := value.(int)
		sparkLineOptions := options.([]sparkline.Option)

		v.Add([]int{valueInput}, sparkLineOptions...)
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

func newSparklineWidget(label string, color cell.Color) (*sparkline.SparkLine, error) {
	return sparkline.New(
		sparkline.Label(
			label,
			cell.FgColor(color),
		),
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

func DashboardLoadingBlinkValue() string {
	inProgress := strings.Repeat("-", 50)
	second := time.Now().Second() % 10

	if second%2 == 0 {
		inProgress = strings.Repeat("\u0020", 50)
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
		case CellNameEpochProgress:
			color = cell.ColorGreen
		case CellNameDiskUsage:
			color = cell.ColorBlue
		case CellNameCpuUsage:
			color = cell.ColorYellow
		case CellNameMemoryUsage:
			color = cell.ColorRed
		}

		if widget, err = newDonutWidget(color); err == nil {
			widget.Percent(1, donut.Label("LOADING...", cell.FgColor(cell.ColorWhite), cell.Bold()))

			return widget
		}
	case CellNameTPSTracker, CellNameCPSTracker:
		var (
			widget *sparkline.SparkLine
			color  cell.Color
		)

		switch name {
		case CellNameTPSTracker:
			color = cell.ColorGreen
		case CellNameCPSTracker:
			color = cell.ColorBlue
		}

		if widget, err = newSparklineWidget("", color); err == nil {
			return widget
		}
	default:
		var widget *segmentdisplay.SegmentDisplay

		if widget, err = newDisplayWidget(); err == nil {
			widget.Write([]*segmentdisplay.TextChunk{
				segmentdisplay.NewChunk(DashboardLoadingBlinkValue(), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite))),
			})

			return widget
		}
	}

	panic(err)
}
