package dashboards

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/container/grid"
	"github.com/mum4k/termdash/widgetapi"
	"github.com/mum4k/termdash/widgets/gauge"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/utility"
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
// It accepts a value to write.
// The type of value must match the type expected by the cell widget.
// If the widget type is not recognized, the function returns nil.
func (c *Cell) Write(value any) error {
	switch widget := c.Widget.(type) {
	case *text.Text:
		return writeToTextWidget(widget, value)
	case *gauge.Gauge:
		return writeToGaugeWidget(widget, value)
	case *segmentdisplay.SegmentDisplay:
		return writeToSegmentWidget(widget, value)
	}

	return nil
}

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

// writeToSegmentWidget writes a value to a segment display widget.
// It accepts a segment display widget and a value to write.
// The value can be an integer, a string, or a slice of strings.
// If the value is a string and it is empty, a blinking value will be used.
// If a string in the slice is empty, a blinking value will be used for that chunk.
func writeToSegmentWidget(widget *segmentdisplay.SegmentDisplay, value any) error {
	capacity := widget.Capacity()

	var chunks []*segmentdisplay.TextChunk

	switch v := value.(type) {
	case int:
		chunk := strconv.Itoa(v)

		chunk = centerOnDisplay(chunk, capacity)

		chunks = append(chunks, segmentdisplay.NewChunk(chunk))
	case string:
		chunk := v

		if chunk == "" {
			chunk = dashboardLoadingBlinkValue(capacity)
		} else {
			chunk = centerOnDisplay(chunk, capacity)
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

// centerOnDisplay takes a string and a number of characters and returns a centered string with empty spaces added at the end
func centerOnDisplay(s string, numChars int) string {
	if len(s) >= numChars {
		return s
	} else {
		// Calculate the number of empty spaces to add on each side of the string
		numSpaces := (numChars - len(s)) / 2
		// Create the empty spaces string
		spaces := strings.Repeat(" ", numSpaces)
		// Concatenate the empty spaces and the original string
		centeredString := spaces + s + spaces
		// If the number of characters is odd, add one more space to the end
		if (numChars - len(centeredString)) == 1 {
			centeredString += " "
		}

		return centeredString
	}
}
