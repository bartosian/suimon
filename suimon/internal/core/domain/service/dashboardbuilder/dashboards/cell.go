package dashboards

import (
	"fmt"
	"strconv"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

// CellsConfig is a type that represents a mapping of column names to strings.
type CellsConfig map[enums.ColumnName]string

// Cells is a type that represents a mapping of column names to pointers to Cell structs.
type Cells map[enums.ColumnName]*Cell

// Cell is a struct that represents a single cell in a dashboard grid. It contains a widget and a list of options.
type Cell struct {
	Widget  widgetapi.Widget
	Options []container.Option
}

// isElement is a method attached to the Cell struct that is used to indicate that it is an Element (part of a grid).
func (Cell) isElement() {}

// NewCell is a function that creates a new Cell struct given a cellName and a widget. It returns a pointer to the new Cell and an error (if any).
func NewCell(cellName string, widget widgetapi.Widget) (*Cell, error) {
	options := append(CellConfigDefault, container.BorderTitle(cellName))

	dashCell := Cell{
		Widget:  widget,
		Options: options,
	}

	return &dashCell, nil
}

// GetWidget is a method attached to the Cell struct that returns a widget that can be added to a dashboard grid.
func (c Cell) GetWidget() grid.Element {
	return grid.Widget(c.Widget, c.Options...)
}

// Write writes a value to the cell widget.
// It accepts a value to write and a set of options.
// The type of value and options must match the type expected by the cell widget.
// If the widget type is not recognized, the function returns nil.
func (c *Cell) Write(value any, options any) error {
	switch widget := c.Widget.(type) {
	case *text.Text:
		return writeToTextWidget(widget, value, options)
	case *gauge.Gauge:
		return writeToGaugeWidget(widget, value, options)
	case *segmentdisplay.SegmentDisplay:
		return writeToSegmentWidget(widget, value, options)
	}

	return nil
}

// writeToTextWidget writes a value to a text widget.
// It accepts a text widget, a value to write, and a set of options.
// The value must be a string type, or an error will be returned.
// Non-printable characters in the value string will be removed before writing to the widget.
// The options must be of type []cell.Option, or an error will be returned.
func writeToTextWidget(widget *text.Text, value any, options any) error {
	valueString, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid value type for text widget: %T", value)
	}

	valueString = log.RemoveNonPrintableChars(valueString)
	if len(valueString) == 0 {
		return nil
	}

	//optionsText, ok := options.([]cell.Option)
	//if !ok {
	//	return fmt.Errorf("invalid options type for text widget: %T", options)
	//}

	return widget.Write(valueString)

	//return widget.Write(valueString, text.WriteCellOpts(optionsText...))
}

// writeToGaugeWidget writes a value to a gauge widget.
// It accepts a gauge widget, a value to write, and a set of options.
// The value must be an integer type, or an error will be returned.
// The options must be of type []gauge.Option, or an error will be returned.
func writeToGaugeWidget(widget *gauge.Gauge, value any, options any) error {
	valueInt, ok := value.(int)
	if !ok {
		return fmt.Errorf("invalid value type for gauge widget: %T", value)
	}

	//optionsGauge, ok := options.([]gauge.Option)
	//if !ok {
	//	return fmt.Errorf("invalid options type for gauge widget: %T", options)
	//}

	return widget.Percent(valueInt)

	//return widget.Percent(valueInt, optionsGauge...)
}

// writeToSegmentWidget writes a value to a segment display widget.
// It accepts a segment display widget, a value to write, and a set of options.
// The options must be of type []segmentdisplay.WriteOption, or an error will be returned.
func writeToSegmentWidget(widget *segmentdisplay.SegmentDisplay, value any, options any) error {
	var segments []*segmentdisplay.TextChunk

	//optionsSegment, ok := options.([]segmentdisplay.WriteOption)
	//if !ok {
	//	return fmt.Errorf("invalid options type for segment widget: %T", options)
	//}

	switch v := value.(type) {
	case int:
		chunk := strconv.Itoa(v)

		segments = append(segments, segmentdisplay.NewChunk(chunk))

		//segments = append(segments, segmentdisplay.NewChunk(chunk, optionsSegment...))
	case string:
		if v == "" {
			v = DashboardLoadingBlinkValue()
		}

		segments = append(segments, segmentdisplay.NewChunk(v))

		//segments = append(segments, segmentdisplay.NewChunk(v, optionsSegment...))
	case []string:
		for _, chunk := range v {
			if chunk == "" {
				chunk = DashboardLoadingBlinkValue()
			}

			segments = append(segments, segmentdisplay.NewChunk(chunk))

			//segments = append(segments, segmentdisplay.NewChunk(chunk, optionsSegment[idx]))
		}
	}

	return widget.Write(segments)
}
