package tablebuilder

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
)

type (
	Builders struct {
		peerTable             Builder
		nodeTable             Builder
		validatorTable        Builder
		rpcTable              Builder
		systemStateTable      Builder
		validatorCountsTable  Builder
		atRiskValidatorsTable Builder
		validatorReportsTable Builder
		activeValidatorsTable Builder
	}

	Builder struct {
		cliGateway *cligw.Gateway
		builder    table.Writer
		config     *TableConfig
	}
)

// NewBuilder creates a new instance of the table builder, using the CLI gateway
func NewBuilder(cliGateway *cligw.Gateway) *Builder {
	tableWR := table.NewWriter()
	tableWR.SetOutputMirror(os.Stdout)

	return &Builder{
		cliGateway: cliGateway,
		builder:    tableWR,
	}
}

// setColumns sets the column configurations for the table builder based on the configuration in the builder's table config
func (tb *Builder) setColumns() {
	var columnsConfig []table.ColumnConfig

	for _, column := range tb.config.Columns {
		columnsConfig = append(columnsConfig, *column.Config)
	}

	tb.builder.SetColumnConfigs(columnsConfig)
}

// setRows sets the rows of the table builder based on the configuration in the builder's table config
func (tb *Builder) setRows() {
	rowsConfig := tb.config.Rows
	columnsConfig := tb.config.Columns
	itemsCount := tb.config.RowsCount
	columnsPerRow := len(rowsConfig[0])

	for itemIndex := 0; itemIndex < itemsCount; itemIndex++ {
		for rowIndex, columns := range rowsConfig {
			header := NewRow(true, false, columnsPerRow, true, text.AlignCenter)
			footer := NewRow(false, false, columnsPerRow, true, text.AlignCenter)
			row := NewRow(false, true, columnsPerRow, true, text.AlignCenter)

			var (
				columnIdx  int
				columnName enums.ColumnName
			)

			for columnIdx, columnName = range columns {
				columnConfig := columnsConfig[columnName]
				columnValue := columnConfig.Values[itemIndex]

				header.AppendValue(columnName.ToString())
				row.AppendValue(columnValue)
				footer.PrependValue(EmptyValue)
			}

			columnIdx++

			for columnIdx < columnsPerRow {
				header.PrependValue(EmptyValue)
				footer.PrependValue(EmptyValue)
				row.PrependValue(EmptyValue)

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

// setStyle sets the style for the table builder based on the configuration in the builder's table config
func (tb *Builder) setStyle() {
	tb.builder.SortBy(tb.config.Sort)
	tb.builder.SetTitle(tb.config.Name)
	tb.builder.SetStyle(tb.config.Style)
	tb.builder.SetAutoIndex(tb.config.AutoIndex)

	tb.setColors()
}

// setColors sets the row colors for the table builder based on the current state of the table
func (tb *Builder) setColors() {
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

func (tb *Builder) Render() {
	tb.setRows()
	tb.setColumns()
	tb.setStyle()

	tb.builder.Render()
}
