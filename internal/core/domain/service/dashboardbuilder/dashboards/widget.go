package dashboards

import (
	"fmt"

	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
)

const gaugeHeight4 = 4
const gaugeTreshold99 = 99
const maxLenBlinkValue = 50

// newWidgetOfType initializes a new widget of the given type.
// It returns the new widget and an error, if any.
func newWidgetOfType(widgetType enums.WidgetType, color cell.Color) (widgetapi.Widget, error) {
	switch widgetType {
	case enums.WidgetTypeProgress:
		return newProgressWidget(color)
	case enums.WidgetTypeTextNoScroll:
		return newTextNoScrollWidget()
	case enums.WidgetTypeDisplay:
		return newDisplayWidget()
	case enums.WidgetTypeSparkLine:
		return newSparklineWidget(color)
	default:
		return nil, fmt.Errorf("invalid widget type: %d", widgetType)
	}
}

// newWidgetByColumnName initializes a new widget based on the given column name.
// It returns the new widget and an error, if any.
func newWidgetByColumnName(columnName enums.ColumnName, color cell.Color) (widgetapi.Widget, error) {
	//nolint: exhaustive,gocritic // no need to cover all cases here.
	switch columnName {
	case enums.ColumnNameTXSyncPercentage, enums.ColumnNameCheckSyncPercentage:
		widget, err := newWidgetOfType(enums.WidgetTypeProgress, color)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize gauge widget for %s: %w", columnName, err)
		}

		progressWidget, ok := widget.(*gauge.Gauge)
		if !ok {
			return nil, fmt.Errorf("failed to cast widget type")
		}

		if err = progressWidget.Percent(0); err != nil {
			return nil, fmt.Errorf("failed to set initial value for %s: %w", columnName, err)
		}

		return widget, nil
	case enums.ColumnNameHealth:
		widget, err := newWidgetOfType(enums.WidgetTypeTextNoScroll, color)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize text widget for %s: %w", columnName, err)
		}

		textWidget, ok := widget.(*text.Text)
		if !ok {
			return nil, fmt.Errorf("failed to cast widget type")
		}

		if err = textWidget.Write(enums.StatusGrey.DashboardStatus(), text.WriteCellOpts(cell.FgColor(cell.ColorGray), cell.BgColor(cell.ColorGray))); err != nil {
			return nil, fmt.Errorf("failed to set initial value for %s: %w", columnName, err)
		}

		return widget, nil
	case enums.ColumnNameCheckpointsPerSecond, enums.ColumnNameTransactionsPerSecond, enums.ColumnNameRoundsPerSecond, enums.ColumnNameCertificatesPerSecond:
		widget, err := newWidgetOfType(enums.WidgetTypeSparkLine, color)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize text widget for %s: %w", columnName, err)
		}

		sparkLineWidget, ok := widget.(*sparkline.SparkLine)
		if !ok {
			return nil, fmt.Errorf("failed to cast widget type")
		}

		if err = sparkLineWidget.Add([]int{0}); err != nil {
			return nil, fmt.Errorf("failed to set initial value for %s: %w", columnName, err)
		}

		return widget, nil
	default:
		widget, err := newWidgetOfType(enums.WidgetTypeDisplay, color)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize segment display widget for %s: %w", columnName, err)
		}

		displayWidget, ok := widget.(*segmentdisplay.SegmentDisplay)
		if !ok {
			return nil, fmt.Errorf("failed to cast widget type")
		}

		err = displayWidget.Write([]*segmentdisplay.TextChunk{
			segmentdisplay.NewChunk(dashboardLoadingBlinkValue(maxLenBlinkValue), segmentdisplay.WriteCellOpts(cell.FgColor(cell.ColorWhite))),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to set initial value for %s: %w", columnName, err)
		}

		return widget, nil
	}
}

// newProgressWidget initializes a new progress widget with the given options.
// It returns the new widget and an error, if any.
func newProgressWidget(color cell.Color) (*gauge.Gauge, error) {
	return gauge.New(
		gauge.Height(gaugeHeight4),
		gauge.Border(linestyle.Light, cell.FgColor(color)),
		gauge.Color(color),
		gauge.FilledTextColor(cell.ColorBlack),
		gauge.EmptyTextColor(cell.ColorWhite),
		gauge.HorizontalTextAlign(align.HorizontalCenter),
		gauge.VerticalTextAlign(align.VerticalMiddle),
		gauge.Threshold(gaugeTreshold99, linestyle.Double, cell.FgColor(cell.ColorRed), cell.Bold()),
	)
}

// newTextNoScrollWidget initializes a new text widget that disables scrolling and wraps at rune boundaries.
// It returns the new widget and an error, if any.
func newTextNoScrollWidget() (*text.Text, error) {
	return text.New(text.DisableScrolling(), text.WrapAtRunes())
}

// newSparklineWidget initializes a new sparkline widget with the given label and color.
// It returns the new widget and an error, if any.
func newSparklineWidget(color cell.Color) (*sparkline.SparkLine, error) {
	return sparkline.New(sparkline.Color(color))
}

// newDisplayWidget initializes a new segment display widget with default options.
// It returns the new widget and an error, if any.
func newDisplayWidget() (*segmentdisplay.SegmentDisplay, error) {
	return segmentdisplay.New(
		segmentdisplay.AlignHorizontal(align.HorizontalCenter),
		segmentdisplay.AlignVertical(align.VerticalMiddle),
		segmentdisplay.MaximizeDisplayedText(),
	)
}
