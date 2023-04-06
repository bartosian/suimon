package tablebuilder

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
)

type (
	TableBuilder struct {
		builder table.Writer
		config  TableConfig
	}
	TableConfig struct {
		Name         string
		Style        table.Style
		SortConfig   []table.SortBy
		AutoIndex    bool
		Columns      map[enums.ColumnName]Column
		Rows         [][]enums.ColumnName
		ColumnsCount int
		RowsCount    int
	}
)

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
	rowsConfig := tb.config.Rows
	columnsConfig := tb.config.Columns
	itemsCount := tb.config.RowsCount

	for itemIndex := 0; itemIndex < itemsCount; itemIndex++ {
		colsPerRow := len(rowsConfig[0])

		for rowIndex, columns := range rowsConfig {
			header := NewRow(true, colsPerRow)
			footer := NewRow(false, colsPerRow)
			row := NewRow(false, colsPerRow)

			var (
				columnIdx  int
				columnName enums.ColumnName
			)

			for columnIdx, columnName = range columns {
				columnConfig := columnsConfig[columnName]
				columnValue := columnConfig.Values[itemIndex]

				header.AppendValue(string(columnName))
				row.AppendValue(columnValue)
				footer.PrependValue("")
			}

			columnIdx++

			for columnIdx < colsPerRow {
				header.PrependValue("")
				footer.PrependValue("")
				row.PrependValue("")

				columnIdx++
			}

			if itemIndex == 0 && rowIndex == 0 {
				tb.builder.AppendHeader(header.Values, header.Config)
				tb.builder.AppendFooter(footer.Values, footer.Config)
			} else if rowIndex%2 == 1 || itemIndex > 0 && len(rowsConfig) > 1 && rowIndex%2 == 0 {
				tb.builder.AppendRow(header.Values, header.Config)
			}

			tb.builder.AppendRow(row.Values, row.Config)
			tb.builder.AppendSeparator()
		}
	}
}

func (tb *TableBuilder) SetStyle() {
	tb.builder.SortBy(tb.config.SortConfig)
	tb.builder.SetTitle(tb.config.Name)
	tb.builder.SetStyle(tb.config.Style)
	tb.builder.SetAutoIndex(tb.config.AutoIndex)

	tb.SetColors()
}

func (tb *TableBuilder) SetColors() {
	var f = func() func(row table.Row) text.Colors {
		valuesRowFgColor := text.Colors{text.FgWhite}
		bgColor := []text.Color{text.BgWhite, text.BgHiBlue, text.BgHiBlue, text.BgWhite}
		currentColor := 0

		var handler = func(row table.Row) text.Colors {
			for _, column := range row {
				switch value := column.(type) {
				case int:
					return valuesRowFgColor
				case string:
					if value == EmptyValue {
						return valuesRowFgColor
					}
				}
			}

			colors := text.Colors{text.FgBlack, bgColor[currentColor]}

			currentColor++
			if currentColor > 3 {
				currentColor = 0
			}

			return colors
		}

		return handler
	}()

	tb.builder.SetRowPainter(f)
}

func (tb *TableBuilder) Render() error {
	tb.SetRows()
	tb.SetColumns()
	tb.SetStyle()

	tb.builder.Render()

	return nil
}
