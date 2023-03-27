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

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

const (
	dashboardName     = "💧 SUIMON: PRESS Q or ESC TO QUIT"
	logsWidgetMessage = "Looking for the sui-node process to stream the logs.\n\n"
)

var (
	DashboardConfigNode = []container.Option{
		container.Border(linestyle.Light),
		container.BorderColor(cell.ColorGreen),
		container.BorderTitle(dashboardName),
		container.FocusedColor(cell.ColorGreen),
		container.AlignHorizontal(align.HorizontalCenter),
		container.AlignVertical(align.VerticalMiddle),
		container.TitleColor(cell.ColorRed),
		container.TitleFocusedColor(cell.ColorRed),
		container.Focused(),
	}

	Rows = []grid.Element{
		0: NewRowPct(15,
			Columns[enums.CellNameNodeHealth],
			Columns[enums.CellNameNetworkHealth],
			Columns[enums.CellNameCurrentEpoch],
			Columns[enums.CellNameEpochTimeTillTheEnd],
			Columns[enums.CellNameUptime],
			Columns[enums.CellNameVersion],
		),
		1: NewRowPct(15,
			Columns[enums.CellNameTotalTransactions],
			Columns[enums.CellNameTotalTransactionCertificates],
			Columns[enums.CellNameTotalTransactionEffects],
		),
		2: NewRowPct(15,
			Columns[enums.CellNameLatestCheckpoint],
			Columns[enums.CellNameHighestKnownCheckpoint],
			Columns[enums.CellNameLastExecutedCheckpoint],
			Columns[enums.CellNameHighestSyncedCheckpoint],
		),
		3: NewRowPct(15,
			Columns[enums.CellNameCheckpointSyncBacklog],
			Columns[enums.CellNameCheckpointExecBacklog],
			Columns[enums.CellNameTransactionsPerSecond],
			Columns[enums.CellNameCheckpointsPerSecond],
			NewColumnPct(40,
				Columns[enums.CellNameTXSyncProgress],
				Columns[enums.CellNameCheckSyncProgress],
			),
		),
		4: NewRowPct(12,
			Columns[enums.CellNameBytesSent],
			Columns[enums.CellNameBytesReceived],
			Columns[enums.CellNameTransactionsPerSecond],
			Columns[enums.CellNameCheckpointsPerSecond],
			Columns[enums.CellNameDatabaseSize],
		),
		//3: NewRowPct(50,
		//	NewColumnPct(40,
		//		NewRowPct(50,
		//			Columns[enums.CellNameEpochProgress],
		//			Columns[enums.CellNameDiskUsage],
		//		),
		//		NewRowPct(50,
		//			Columns[enums.CellNameCpuUsage],
		//			Columns[enums.CellNameMemoryUsage],
		//		),
		//	),
		//	NewColumnPct(60,
		//		NewRowPct(15,
		//			Columns[enums.CellNameTXSyncProgress],
		//			Columns[enums.CellNameCheckSyncProgress],
		//		),
		//		NewRowPct(20, Columns[enums.CellNameTPSTracker]),
		//		NewRowPct(20, Columns[enums.CellNameCPSTracker]),
		//		NewRowPct(45, Columns[enums.CellNameNodeLogs]),
		//	),
		//),
	}

	Columns = []grid.Element{
		enums.CellNameNodeHealth:                   NewColumnPct(5, Cells[enums.CellNameNodeHealth].GetGridWidget()),
		enums.CellNameNetworkHealth:                NewColumnPct(5, Cells[enums.CellNameNetworkHealth].GetGridWidget()),
		enums.CellNameCurrentEpoch:                 NewColumnPct(20, Cells[enums.CellNameCurrentEpoch].GetGridWidget()),
		enums.CellNameEpochTimeTillTheEnd:          NewColumnPct(20, Cells[enums.CellNameEpochTimeTillTheEnd].GetGridWidget()),
		enums.CellNameUptime:                       NewColumnPct(20, Cells[enums.CellNameUptime].GetGridWidget()),
		enums.CellNameVersion:                      NewColumnPct(20, Cells[enums.CellNameVersion].GetGridWidget()),
		enums.CellNameCommit:                       NewColumnPct(20, Cells[enums.CellNameCommit].GetGridWidget()),
		enums.CellNameTotalTransactions:            NewColumnPct(30, Cells[enums.CellNameTotalTransactions].GetGridWidget()),
		enums.CellNameTotalTransactionCertificates: NewColumnPct(30, Cells[enums.CellNameTotalTransactionCertificates].GetGridWidget()),
		enums.CellNameTotalTransactionEffects:      NewColumnPct(30, Cells[enums.CellNameTotalTransactionEffects].GetGridWidget()),
		enums.CellNameLatestCheckpoint:             NewColumnPct(25, Cells[enums.CellNameLatestCheckpoint].GetGridWidget()),
		enums.CellNameHighestKnownCheckpoint:       NewColumnPct(25, Cells[enums.CellNameHighestKnownCheckpoint].GetGridWidget()),
		enums.CellNameLastExecutedCheckpoint:       NewColumnPct(25, Cells[enums.CellNameLastExecutedCheckpoint].GetGridWidget()),
		enums.CellNameHighestSyncedCheckpoint:      NewColumnPct(25, Cells[enums.CellNameHighestSyncedCheckpoint].GetGridWidget()),
		enums.CellNameCheckpointSyncBacklog:        NewColumnPct(15, Cells[enums.CellNameCheckpointSyncBacklog].GetGridWidget()),
		enums.CellNameCheckpointExecBacklog:        NewColumnPct(15, Cells[enums.CellNameCheckpointExecBacklog].GetGridWidget()),
		enums.CellNameTransactionsPerSecond:        NewColumnPct(15, Cells[enums.CellNameTransactionsPerSecond].GetGridWidget()),
		enums.CellNameCheckpointsPerSecond:         NewColumnPct(15, Cells[enums.CellNameCheckpointsPerSecond].GetGridWidget()),
		enums.CellNameConnectedPeers:               NewColumnPct(10, Cells[enums.CellNameConnectedPeers].GetGridWidget()),
		enums.CellNameTXSyncProgress:               NewColumnPct(50, Cells[enums.CellNameTXSyncProgress].GetGridWidget()),
		enums.CellNameCheckSyncProgress:            NewColumnPct(50, Cells[enums.CellNameCheckSyncProgress].GetGridWidget()),
		enums.CellNameEpochProgress:                NewColumnPct(50, Cells[enums.CellNameEpochProgress].GetGridWidget()),
		enums.CellNameDiskUsage:                    NewColumnPct(50, Cells[enums.CellNameDiskUsage].GetGridWidget()),
		enums.CellNameDatabaseSize:                 NewColumnPct(20, Cells[enums.CellNameDatabaseSize].GetGridWidget()),
		enums.CellNameBytesSent:                    NewColumnPct(20, Cells[enums.CellNameBytesSent].GetGridWidget()),
		enums.CellNameBytesReceived:                NewColumnPct(20, Cells[enums.CellNameBytesReceived].GetGridWidget()),
		enums.CellNameMemoryUsage:                  NewColumnPct(50, Cells[enums.CellNameMemoryUsage].GetGridWidget()),
		enums.CellNameCpuUsage:                     NewColumnPct(50, Cells[enums.CellNameCpuUsage].GetGridWidget()),
		enums.CellNameNodeLogs:                     NewColumnPct(50, Cells[enums.CellNameNodeLogs].GetGridWidget()),
		enums.CellNameButtonQuit:                   NewColumnPct(25, Cells[enums.CellNameButtonQuit].GetGridWidget()),
		enums.CellNameTPSTracker:                   NewColumnPct(99, Cells[enums.CellNameTPSTracker].GetGridWidget()),
		enums.CellNameCPSTracker:                   NewColumnPct(99, Cells[enums.CellNameCPSTracker].GetGridWidget()),
	}

	Cells = []*Cell{
		enums.CellNameNodeHealth:                   NewCell(enums.CellNameNodeHealth),
		enums.CellNameNetworkHealth:                NewCell(enums.CellNameNetworkHealth),
		enums.CellNameTransactionsPerSecond:        NewCell(enums.CellNameTransactionsPerSecond),
		enums.CellNameCheckpointsPerSecond:         NewCell(enums.CellNameCheckpointsPerSecond),
		enums.CellNameTotalTransactions:            NewCell(enums.CellNameTotalTransactions),
		enums.CellNameTotalTransactionCertificates: NewCell(enums.CellNameTotalTransactionCertificates),
		enums.CellNameTotalTransactionEffects:      NewCell(enums.CellNameTotalTransactionEffects),
		enums.CellNameLatestCheckpoint:             NewCell(enums.CellNameLatestCheckpoint),
		enums.CellNameCheckpointSyncBacklog:        NewCell(enums.CellNameCheckpointSyncBacklog),
		enums.CellNameCheckpointExecBacklog:        NewCell(enums.CellNameCheckpointExecBacklog),
		enums.CellNameHighestKnownCheckpoint:       NewCell(enums.CellNameHighestKnownCheckpoint),
		enums.CellNameLastExecutedCheckpoint:       NewCell(enums.CellNameLastExecutedCheckpoint),
		enums.CellNameHighestSyncedCheckpoint:      NewCell(enums.CellNameHighestSyncedCheckpoint),
		enums.CellNameConnectedPeers:               NewCell(enums.CellNameConnectedPeers),
		enums.CellNameTXSyncProgress:               NewCell(enums.CellNameTXSyncProgress),
		enums.CellNameCheckSyncProgress:            NewCell(enums.CellNameCheckSyncProgress),
		enums.CellNameUptime:                       NewCell(enums.CellNameUptime),
		enums.CellNameVersion:                      NewCell(enums.CellNameVersion),
		enums.CellNameCommit:                       NewCell(enums.CellNameCommit),
		enums.CellNameCurrentEpoch:                 NewCell(enums.CellNameCurrentEpoch),
		enums.CellNameEpochProgress:                NewCell(enums.CellNameEpochProgress),
		enums.CellNameEpochTimeTillTheEnd:          NewCell(enums.CellNameEpochTimeTillTheEnd),
		enums.CellNameDiskUsage:                    NewCell(enums.CellNameDiskUsage),
		enums.CellNameDatabaseSize:                 NewCell(enums.CellNameDatabaseSize),
		enums.CellNameBytesSent:                    NewCell(enums.CellNameBytesSent),
		enums.CellNameBytesReceived:                NewCell(enums.CellNameBytesReceived),
		enums.CellNameMemoryUsage:                  NewCell(enums.CellNameMemoryUsage),
		enums.CellNameCpuUsage:                     NewCell(enums.CellNameCpuUsage),
		enums.CellNameNodeLogs:                     NewCell(enums.CellNameNodeLogs),
		enums.CellNameButtonQuit:                   NewCell(enums.CellNameButtonQuit),
		enums.CellNameTPSTracker:                   NewCell(enums.CellNameTPSTracker),
		enums.CellNameCPSTracker:                   NewCell(enums.CellNameCPSTracker),
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
		Name    enums.CellName
		Widget  widgetapi.Widget
		Options []container.Option
	}
)

func (Row) isElement() {}

func (Column) isElement() {}

func (Cell) isElement() {}

func NewCell(name enums.CellName) *Cell {
	return &Cell{
		Name:   name,
		Widget: newWidgetByCellName(name),
		Options: []container.Option{
			container.Border(linestyle.Light),
			container.BorderTitle(name.CellNameString()),
			container.BorderColor(cell.ColorGreen),
			container.AlignVertical(align.VerticalMiddle),
			container.AlignHorizontal(align.HorizontalCenter),
			container.TitleColor(cell.ColorRed),
			container.FocusedColor(cell.ColorWhite),
			container.FocusedColor(cell.ColorBlue),
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

		if c.Name != enums.CellNameNodeLogs {
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
		gauge.TextLabel("RPC delta"),
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
		segmentdisplay.AlignHorizontal(align.HorizontalCenter),
		segmentdisplay.AlignVertical(align.VerticalMiddle),
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

func newWidgetByCellName(name enums.CellName) widgetapi.Widget {
	var err error

	switch name {
	case enums.CellNameCheckSyncProgress, enums.CellNameTXSyncProgress:
		var widget *gauge.Gauge

		if widget, err = newProgressWidget(); err == nil {
			widget.Percent(0)

			return widget
		}
	case enums.CellNameNodeHealth, enums.CellNameNetworkHealth:
		var widget *text.Text

		if widget, err = newTextNoScrollWidget(); err == nil {
			widget.Write(enums.StatusGrey.DashboardStatus(), text.WriteCellOpts(cell.FgColor(cell.ColorGray), cell.BgColor(cell.ColorGray)))

			return widget
		}
	case enums.CellNameNodeLogs:
		var widget *text.Text

		if widget, err = newTextWidget(); err == nil {
			widget.Write(logsWidgetMessage, text.WriteCellOpts(cell.FgColor(cell.ColorWhite), cell.Bold()))

			return widget
		}
	case enums.CellNameEpochProgress, enums.CellNameDiskUsage, enums.CellNameMemoryUsage, enums.CellNameCpuUsage:
		var (
			widget *donut.Donut
			color  cell.Color
		)

		switch name {
		case enums.CellNameEpochProgress:
			color = cell.ColorGreen
		case enums.CellNameDiskUsage:
			color = cell.ColorBlue
		case enums.CellNameCpuUsage:
			color = cell.ColorYellow
		case enums.CellNameMemoryUsage:
			color = cell.ColorRed
		}

		if widget, err = newDonutWidget(color); err == nil {
			widget.Percent(1, donut.Label("LOADING...", cell.FgColor(cell.ColorWhite), cell.Bold()))

			return widget
		}
	case enums.CellNameTPSTracker, enums.CellNameCPSTracker:
		var (
			widget *sparkline.SparkLine
			color  cell.Color
		)

		switch name {
		case enums.CellNameTPSTracker:
			color = cell.ColorGreen
		case enums.CellNameCPSTracker:
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
