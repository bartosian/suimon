package tablebuilder

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type TableBuilder struct {
	builder table.Writer
	config  TableConfig
}

type TableConfig struct {
	Name         string
	Style        table.Style
	Columns      []Column
	SortConfig   []table.SortBy
	RowsCount    int
	ColumnsCount int
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

	for i := 0; i < tb.config.RowsCount; i++ {
		rowValues := make([]any, 0, tb.config.ColumnsCount)

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
	tb.builder.Style().Title.Colors = text.Colors{text.FgBlue}
	tb.builder.Style().Title.Align = text.AlignCenter
	tb.builder.Style().Box.RightSeparator = ""
	tb.builder.SetAutoIndex(true)
	tb.builder.Style().Color = table.ColorOptions{
		Header: text.Colors{text.FgHiRed},
		Row:    text.Colors{text.FgHiWhite},
		Footer: text.Colors{text.FgBlue},
	}
	tb.builder.Style().Color.IndexColumn = text.Colors{text.FgHiRed}
}

func (tb *TableBuilder) Build() {
	tb.SetRows()
	tb.SetColumns()
	tb.SetStyle()

	tb.builder.Render()
}
