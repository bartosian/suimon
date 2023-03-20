package tablebuilder

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/tablebuilder/tables"
)

type (
	TableBuilder struct {
		builder table.Writer
		config  TableConfig
	}
	TableColors struct {
		Title  text.Color
		Header text.Color
		Row    text.Color
		Footer text.Color
		Index  text.Color
	}
	TableConfig struct {
		Name         string
		Tag          string
		Colors       TableColors
		Style        table.Style
		Columns      []Column
		SortConfig   []table.SortBy
		RowsCount    int
		ColumnsCount int
	}
)

var tbColorsValues = map[enums.ColorTable]TableColors{
	enums.ColorTableWhite: {Title: text.FgHiWhite, Header: text.FgHiWhite, Row: text.FgHiWhite, Footer: text.FgHiWhite, Index: text.FgHiWhite},
	enums.ColorTableDark:  {Title: text.FgBlack, Header: text.FgBlack, Row: text.FgBlack, Footer: text.FgBlack, Index: text.FgBlack},
	enums.ColorTableColor: {Title: text.FgHiBlue, Header: text.FgHiRed, Row: text.FgHiWhite, Footer: text.FgHiBlue, Index: text.FgHiRed},
}

// GetTableColorsFromString converts a color string to a TableColors struct.
// Returns: If the color string is invalid, GetTableColorsFromString returns the default colors.
func GetTableColorsFromString(color string) TableColors {
	colorTable := enums.ColorTableFromString(color)

	return tbColorsValues[colorTable]
}

func NewTableBuilder(config TableConfig) *TableBuilder {
	tableWR := table.NewWriter()
	tableWR.SetOutputMirror(os.Stdout)

	return &TableBuilder{
		builder: tableWR,
		config:  config,
	}
}

func (tb *TableBuilder) SetColumns() {
	var columnsConfig []table.ColumnConfig

	for _, column := range tb.config.Columns {
		columnsConfig = append(columnsConfig, column.Config)
	}

	tb.builder.SetColumnConfigs(columnsConfig)
}

func (tb *TableBuilder) SetRows() {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}

	var (
		header table.Row
		footer table.Row
	)

	for _, column := range tb.config.Columns {
		header = append(header, column.Config.Name)
		footer = append(footer, "")
	}

	footer[0] = fmt.Sprintf("%s%s%s", enums.ColorRed, tb.config.Tag, enums.ColorReset)

	for i := 0; i < tb.config.RowsCount; i++ {
		rowValues := make([]any, 0, len(tables.ColumnConfigNode))

		for _, column := range tb.config.Columns {
			rowValues = append(rowValues, column.Values[i])
		}

		tb.builder.AppendRow(rowValues, rowConfigAutoMerge)
		tb.builder.AppendSeparator()
	}

	tb.builder.AppendHeader(header, rowConfigAutoMerge)
	tb.builder.AppendFooter(footer, rowConfigAutoMerge)
	tb.builder.SortBy(tb.config.SortConfig)
}

func (tb *TableBuilder) SetStyle() {
	tb.builder.SetTitle(tb.config.Name)
	tb.builder.SetStyle(tb.config.Style)
	tb.builder.Style().Title.Align = text.AlignLeft
	tb.builder.Style().Box.RightSeparator = ""
	tb.builder.SetAutoIndex(true)
	tb.SetColors()
}

func (tb *TableBuilder) SetColors() {
	colors := tb.config.Colors

	tb.builder.Style().Title.Colors = text.Colors{colors.Title}
	tb.builder.Style().Color = table.ColorOptions{
		Header: text.Colors{colors.Header},
		Row:    text.Colors{colors.Row},
		Footer: text.Colors{colors.Footer},
	}
	tb.builder.Style().Color.IndexColumn = text.Colors{colors.Index}
}

func (tb *TableBuilder) Build() {
	tb.SetRows()
	tb.SetColumns()
	tb.SetStyle()

	tb.builder.Render()
}
