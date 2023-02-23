package dashboards

import (
	"fmt"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
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
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
)

const suiEmoji = "💧"

var (
	DashboardConfigSUI = []container.Option{
		container.Border(linestyle.Light),
		container.BorderColor(cell.ColorGreen),
		container.FocusedColor(cell.ColorGreen),
		container.AlignHorizontal(align.HorizontalCenter),
		container.AlignVertical(align.VerticalMiddle),
		container.BorderTitle(fmt.Sprintf("%s SUIMON: PRESS Q or ESC TO QUIT", suiEmoji)),
		container.Focused(),
	}

	Rows = []grid.Element{
		0: NewRow(7,
			Columns[CellNameNodeStatus],
			Columns[CellNameNetworkStatus],
			Columns[CellNameTotalTransactions],
			Columns[CellNameLatestCheckpoint],
			Columns[CellNameHighestCheckpoint],
			NewColumn(0), // window width limiter
		),
		1: NewRow(7,
			Columns[CellNameUptime],
			Columns[CellNameTransactionsPerSecond],
			Columns[CellNameCheckpointsPerSecond],
			Columns[CellNameConnectedPeers],
			Columns[CellNameVersion],
			Columns[CellNameCommit],
			NewColumn(0), // window width limiter
		),
		2: NewRow(7,
			Columns[CellNameEpochEnd],
			Columns[CellNameDatabaseSize],
			Columns[CellNameTXSyncProgress],
			Columns[CellNameCheckSyncProgress],
			NewColumn(0), // window width limiter
		),
		3: NewRow(42,
			NewColumn(86,
				NewRow(21, Columns[CellNameEpoch], Columns[CellNameDiskUsage]),
				NewRow(21, Columns[CellNameCpuUsage], Columns[CellNameMemoryUsage])),
			NewColumn(140, Columns[CellNameNodeLogs]),
			NewColumn(0), // window width limiter
		),
		4: NewRow(0), // window height limiter
	}

	Columns = []grid.Element{
		CellNameNodeStatus:            NewColumn(8, Cells[CellNameNodeStatus].GetGridWidget()),
		CellNameNetworkStatus:         NewColumn(8, Cells[CellNameNetworkStatus].GetGridWidget()),
		CellNameAddress:               NewColumn(8, Cells[CellNameAddress].GetGridWidget()),
		CellNameTransactionsPerSecond: NewColumn(28, Cells[CellNameTransactionsPerSecond].GetGridWidget()),
		CellNameCheckpointsPerSecond:  NewColumn(26, Cells[CellNameCheckpointsPerSecond].GetGridWidget()),
		CellNameTotalTransactions:     NewColumn(70, Cells[CellNameTotalTransactions].GetGridWidget()),
		CellNameLatestCheckpoint:      NewColumn(70, Cells[CellNameLatestCheckpoint].GetGridWidget()),
		CellNameHighestCheckpoint:     NewColumn(70, Cells[CellNameHighestCheckpoint].GetGridWidget()),
		CellNameConnectedPeers:        NewColumn(25, Cells[CellNameConnectedPeers].GetGridWidget()),
		CellNameTXSyncProgress:        NewColumn(70, Cells[CellNameTXSyncProgress].GetGridWidget()),
		CellNameCheckSyncProgress:     NewColumn(70, Cells[CellNameCheckSyncProgress].GetGridWidget()),
		CellNameUptime:                NewColumn(32, Cells[CellNameUptime].GetGridWidget()),
		CellNameVersion:               NewColumn(45, Cells[CellNameVersion].GetGridWidget()),
		CellNameCommit:                NewColumn(70, Cells[CellNameCommit].GetGridWidget()),
		CellNameCompany:               NewColumn(100, Cells[CellNameCompany].GetGridWidget()),
		CellNameCountry:               NewColumn(100, Cells[CellNameCountry].GetGridWidget()),
		CellNameEpoch:                 NewColumn(43, Cells[CellNameEpoch].GetGridWidget()),
		CellNameEpochEnd:              NewColumn(43, Cells[CellNameEpochEnd].GetGridWidget()),
		CellNameDiskUsage:             NewColumn(43, Cells[CellNameDiskUsage].GetGridWidget()),
		CellNameDatabaseSize:          NewColumn(43, Cells[CellNameDatabaseSize].GetGridWidget()),
		CellNameBytesSent:             NewColumn(43, Cells[CellNameBytesSent].GetGridWidget()),
		CellNameBytesReceived:         NewColumn(43, Cells[CellNameBytesReceived].GetGridWidget()),
		CellNameMemoryUsage:           NewColumn(43, Cells[CellNameMemoryUsage].GetGridWidget()),
		CellNameCpuUsage:              NewColumn(43, Cells[CellNameCpuUsage].GetGridWidget()),
		CellNameNodeLogs:              NewColumn(140, Cells[CellNameNodeLogs].GetGridWidget()),
	}

	Cells = []*Cell{
		CellNameNodeStatus:            NewCell("NODE", CellNameNodeStatus),
		CellNameNetworkStatus:         NewCell("NET", CellNameNetworkStatus),
		CellNameAddress:               NewCell("ADDRESS", CellNameAddress),
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
		CellNameCompany:               NewCell("PROVIDER", CellNameCompany),
		CellNameCountry:               NewCell("COUNTRY", CellNameCountry),
		CellNameEpoch:                 NewCell("EPOCH", CellNameEpoch),
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

type Element interface {
	isElement()
}

type Row struct {
	Height   int
	Elements []Element
}

func (Row) isElement() {}

type Column struct {
	Width    int
	Elements []Element
}

func (Column) isElement() {}

type Cell struct {
	Name    CellName
	Widget  widgetapi.Widget
	Options []container.Option
}

func (Cell) isElement() {}

func NewCell(title string, name CellName) *Cell {
	return &Cell{
		Name:   name,
		Widget: newWidgetByCellName(name),
		Options: []container.Option{
			container.FocusedColor(cell.ColorGreen),
			container.Border(linestyle.Light),
			container.BorderTitle(title),
			container.AlignVertical(align.VerticalMiddle),
			container.AlignHorizontal(align.HorizontalCenter),
			container.BorderColor(cell.ColorRed),
			container.TitleColor(cell.ColorGreen),
		},
	}
}

func NewRow(height int, elements ...grid.Element) grid.Element {
	return grid.RowHeightFixed(height, elements...)
}

func NewColumn(width int, elements ...grid.Element) grid.Element {
	return grid.ColWidthFixed(width, elements...)
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

func (c *Cell) Write(value any, options ...cell.Option) {
	switch v := c.Widget.(type) {
	case *text.Text:
		valueString := value.(string)

		if c.Name != CellNameNodeLogs {
			v.Reset()
		}

		v.Write(valueString, text.WriteCellOpts(options...))
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
			chunkValue = dashboardLoadingBlinkValue()
		}

		v.Write([]*segmentdisplay.TextChunk{
			segmentdisplay.NewChunk(chunkValue, segmentdisplay.WriteCellOpts(options...)),
		})
	case *donut.Donut:
		valueInput := value.(DonutWriteInput)

		v.Percent(
			valueInput.Percentage,
			donut.Label(valueInput.Label, options...),
		)
	}
}

func (c *Cell) GetGridWidget() grid.Element {
	return grid.Widget(c.Widget, c.Options...)
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

func newTextWidget() (*text.Text, error) {
	return text.New(text.RollContent(), text.WrapAtWords())
}

func newDonutWidget(color cell.Color) (*donut.Donut, error) {
	return donut.New(
		donut.CellOpts(
			cell.FgColor(color),
			cell.Bold(),
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

		if widget, err = newTextWidget(); err == nil {
			widget.Write(enums.StatusGrey.DashboardStatus(), text.WriteCellOpts(cell.FgColor(cell.ColorGray), cell.BgColor(cell.ColorGray)))

			return widget
		}
	case CellNameNodeLogs:
		var widget *text.Text

		if widget, err = newTextWidget(); err == nil {
			defaultValue := log.GenerateLogoFrom("  Loading...", "banner", "gray")

			widget.Write("\n\n"+defaultValue, text.WriteCellOpts(cell.FgColor(cell.ColorGray), cell.Bold()))

			return widget
		}
	case CellNameEpoch, CellNameDiskUsage, CellNameMemoryUsage, CellNameCpuUsage:
		var (
			widget *donut.Donut
			color  = cell.ColorGreen
		)

		switch name {
		case CellNameDiskUsage:
			color = cell.ColorBlue
		case CellNameCpuUsage:
			color = cell.ColorYellow
		case CellNameMemoryUsage:
			color = cell.ColorRed
		}

		if widget, err = newDonutWidget(color); err == nil {
			widget.Percent(1, donut.Label("LOADING...", cell.FgColor(cell.ColorGray), cell.Bold()))

			return widget
		}
	default:
		var widget *segmentdisplay.SegmentDisplay

		if widget, err = newDisplayWidget(); err == nil {
			widget.Write([]*segmentdisplay.TextChunk{
				segmentdisplay.NewChunk(dashboardLoadingValue(), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorGray))),
			})

			return widget
		}
	}

	panic(err)
}
