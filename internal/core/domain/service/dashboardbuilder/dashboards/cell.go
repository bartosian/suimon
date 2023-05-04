package dashboards

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/sparkline"
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
)

// CellsConfig is a type that represents a mapping of column names to CellConfig.
type (
	CellsConfig map[enums.ColumnName]CellConfig
	CellConfig  struct {
		Title string
		Color cell.Color
	}
)

// Cells is a type that represents a mapping of column names to pointers to Cell structs.
type Cells map[enums.ColumnName]*Cell

// Cell is a struct that represents a single cell in a dashboard grid. It contains a widget and a list of options.
type Cell struct {
	Widget        widgetapi.Widget
	Options       []container.Option
	LastUpdatedAt time.Time
}

// NewCell is a function that creates a new Cell struct given a cellName and a widget. It returns a pointer to the new Cell and an error (if any).
func NewCell(cellName string, color cell.Color, widget widgetapi.Widget) (*Cell, error) {
	options := append(CellConfigDefault, container.BorderTitle(cellName), container.BorderColor(color))

	dashCell := Cell{
		Widget:        widget,
		Options:       options,
		LastUpdatedAt: time.Now(),
	}

	return &dashCell, nil
}

// GetWidget is a method attached to the Cell struct that returns a widget that can be added to a dashboard grid.
func (c Cell) GetWidget() grid.Element {
	return grid.Widget(c.Widget, c.Options...)
}

// Write writes a value to the cell widget.
// It accepts a value to write.
// The type of value must match the type expected by the cell widget.
// If the widget type is not recognized, the function returns nil.
func (c *Cell) Write(value any) error {
	now := time.Now()

	switch widget := c.Widget.(type) {
	case *gauge.Gauge:
		return writeToGaugeWidget(widget, value)
	case *text.Text:
		return writeToTextWidget(widget, value)
	case *segmentdisplay.SegmentDisplay:
		return writeToSegmentWidget(widget, value)
	case *sparkline.SparkLine:
		if now.Sub(c.LastUpdatedAt) < 2*time.Second {
			return nil
		}

		c.LastUpdatedAt = now

		return writeToSparkLineWidget(widget, value)
	}

	return nil
}

// writeToTextWidget writes a string value to a text widget with the given options.
// The function expects a slice of cell options and a value of type string,
// and returns an error if the value has a different type. The function uses
// the `text.Text` type and its `Write` method to write the string value to the widget,
// with the options converted to `text.WriteOption` using the `text.WriteCellOpts` function.
// The function removes any non-printable characters from the string value before writing it
// to the widget, and returns immediately if the resulting string has zero length.
func writeToTextWidget(widget *text.Text, value any) error {
	valueString, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid value type for text widget: %T", value)
	}

	valueString = log.RemoveNonPrintableChars(valueString)
	if len(valueString) == 0 {
		return nil
	}

	return widget.Write(valueString)
}

// writeToGaugeWidget writes a value to a gauge widget.
// It accepts a gauge widget and a value to write.
// The value must be a string representing an integer, or an error will be returned.
func writeToGaugeWidget(widget *gauge.Gauge, value any) error {
	valueString, ok := value.(string)
	if !ok {
		return fmt.Errorf("unexpected metric value type for gauge widget: %T", value)
	}

	valueInt, err := utility.ParseIntFromString(valueString)
	if err != nil {
		return fmt.Errorf("unexpected metric value type for gauge widget: %T", value)
	}

	return widget.Percent(valueInt)
}

// writeToSparkLineWidget adds a new value to a sparkline chart widget.
// The function expects an integer value and returns an error if the value
// has a different type. It uses the `widget` argument to add the new value
// to the chart using the `Add` method of the `sparkline.SparkLine` type.
func writeToSparkLineWidget(widget *sparkline.SparkLine, value any) error {
	valueInt, ok := value.(int)
	if !ok {
		return fmt.Errorf("unexpected metric value type for sparkline widget: %T", value)
	}

	return widget.Add([]int{valueInt})
}

// writeToSegmentWidget writes a value to a segment display widget.
// It accepts a segment display widget and a value to write.
// The value can be an integer, a string, or a slice of strings.
// If the value is a string and it is empty, a blinking value will be used.
// If a string in the slice is empty, a blinking value will be used for that chunk.
func writeToSegmentWidget(widget *segmentdisplay.SegmentDisplay, value any) error {
	capacity := widget.Capacity()

	var chunks []*segmentdisplay.TextChunk

	switch v := value.(type) {
	case int64:
		chunk := strconv.FormatInt(v, 10)

		chunks = append(chunks, segmentdisplay.NewChunk(chunk))
	case int:
		chunk := strconv.Itoa(v)

		chunks = append(chunks, segmentdisplay.NewChunk(chunk))
	case string:
		chunk := v

		if chunk == "" {
			chunk = dashboardLoadingBlinkValue(capacity)
		}

		chunks = append(chunks, segmentdisplay.NewChunk(chunk))
	case []string:
		for _, chunk := range v {
			if chunk == "" {
				chunk = dashboardLoadingBlinkValue(capacity)
			}

			chunks = append(chunks, segmentdisplay.NewChunk(chunk))
		}
	}

	return widget.Write(chunks)
}

// dashboardLoadingBlinkValue returns a string that represents a loading
// animation that blinks every even second.
//
// The animation consists of a series of "-" characters that fill up the
// capacity parameter, which represents the maximum length of the animation.
// Every even second, the animation is replaced with spaces to create a
// blinking effect.
func dashboardLoadingBlinkValue(maxLength int) string {
	inProgress := strings.Repeat("-", maxLength)
	second := time.Now().Second()

	if second%2 == 0 {
		inProgress = strings.Repeat("\u0020", maxLength)
	}

	return inProgress
}
