package tablebuilder

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/suimon/internal/core/domain/enums"
	"github.com/bartosian/suimon/internal/core/domain/host"
	"github.com/bartosian/suimon/internal/core/domain/metrics"
	"github.com/bartosian/suimon/internal/core/domain/service/tablebuilder/tables"
	"github.com/bartosian/suimon/internal/core/gateways/cligw"
)

const slashingPct50 = 50
const slashingPct100 = 50

type Builder struct {
	writer     table.Writer
	cliGateway *cligw.Gateway
	config     *tables.TableConfig
	tableType  enums.TableType
	hosts      []host.Host
	Releases   []metrics.Release
}

// NewBuilder creates a new instance of the table builder, using the CLI gateway.
func NewBuilder(tableType enums.TableType, hosts []host.Host, releases []metrics.Release, cliGateway *cligw.Gateway) *Builder {
	tableWR := table.NewWriter()
	tableWR.SetOutputMirror(os.Stdout)

	return &Builder{
		tableType:  tableType,
		Releases:   releases,
		hosts:      hosts,
		cliGateway: cliGateway,
		writer:     tableWR,
	}
}

// setColumns sets the column configurations for the table builder based on the configuration in the builder's table config.
func (tb *Builder) setColumns() {
	columnsConfig := make([]table.ColumnConfig, len(tb.config.Columns))

	for _, column := range tb.config.Columns {
		columnsConfig = append(columnsConfig, *column.Config)
	}

	tb.writer.SetColumnConfigs(columnsConfig)
}

// setRows sets the rows of the table builder based on the configuration in the builder's table config.
func (tb *Builder) setRows() error {
	rowsConfig := tb.config.Rows
	columnsConfig := tb.config.Columns
	itemsCount := tb.config.RowsCount
	columnsPerRow := len(rowsConfig[0])

	// Prepare footer with empty values
	footer := tables.NewRow(tables.NewRowConfig{
		IsHeader:       false,
		IsFooter:       true,
		Length:         columnsPerRow,
		AutoMerge:      true,
		AutoMergeAlign: text.AlignCenter,
	})

	for itemIndex := 0; itemIndex < itemsCount; itemIndex++ {
		for rowIndex, columns := range rowsConfig {
			isFirstRow := itemIndex == 0 && rowIndex == 0
			isEvenRow := rowIndex%2 == 0
			multipleRows := len(rowsConfig) > 1

			header := tables.NewRow(tables.NewRowConfig{
				IsHeader:       true,
				IsFooter:       false,
				Length:         columnsPerRow,
				AutoMerge:      true,
				AutoMergeAlign: text.AlignCenter,
			})
			row := tables.NewRow(tables.NewRowConfig{
				IsHeader:       false,
				IsFooter:       true,
				Length:         columnsPerRow,
				AutoMerge:      true,
				AutoMergeAlign: text.AlignCenter,
			})

			// Build the row and header
			for _, columnName := range columns {
				columnConfig, ok := columnsConfig[columnName]
				if !ok {
					tb.cliGateway.Errorf("column %s not found", columnName)
					return fmt.Errorf("column %s not found", columnName)
				}

				columnValue := columnConfig.Values[itemIndex]

				if isFirstRow {
					header.AppendValue(columnName.ToString())
					footer.PrependValue(tables.EmptyValue)
				} else if multipleRows && (rowIndex != 0 || itemIndex > 0 && isEvenRow) {
					header.AppendValue(columnName.ToString())
				}

				row.AppendValue(columnValue)
			}

			// Handle empty spaces in rows
			for i := len(columns); i < columnsPerRow; i++ {
				emptyValue := tables.EmptyValue
				if isFirstRow {
					header.PrependValue(emptyValue)
					footer.PrependValue(emptyValue)
				} else if multipleRows && (rowIndex != 0 || itemIndex > 0 && isEvenRow) {
					header.PrependValue(emptyValue)
				}

				row.PrependValue(emptyValue)
			}

			// Append header and footer for the first row
			if isFirstRow {
				tb.writer.AppendHeader(header.Values, header.Config)
				tb.writer.AppendFooter(footer.Values, footer.Config)
			} else if multipleRows && (rowIndex != 0 || itemIndex > 0 && isEvenRow) {
				tb.writer.AppendRow(header.Values, header.Config)
			}

			// Append the row and separator
			tb.writer.AppendRow(row.Values, row.Config)
			tb.writer.AppendSeparator()
		}
	}

	return nil
}

// setStyle sets the style for the table builder based on the configuration in the builder's table config.
func (tb *Builder) setStyle() {
	tb.writer.SetTitle(tb.config.Name)
	tb.writer.SetStyle(tb.config.Style)

	tb.setColors()
}

// setColors sets the row colors for the table builder based on the current state of the table.
func (tb *Builder) setColors() {
	var (
		fgWhite  = text.FgWhite
		bgWhite  = text.BgWhite
		fgBlack  = text.FgBlack
		bgRed    = text.BgRed
		bgYellow = text.BgYellow
	)

	var painter = func() func(row table.Row) text.Colors {
		valuesRowFgColor := text.Colors{fgWhite}

		var handler = func(row table.Row) text.Colors {
			if tb.tableType == enums.TableTypeValidatorReports {
				valueString, ok := row[1].(string)
				if !ok {
					return valuesRowFgColor
				}

				slashingPct, err := strconv.ParseFloat(valueString, 64)
				if err != nil {
					return valuesRowFgColor
				}

				if slashingPct > slashingPct100 {
					return text.Colors{bgRed, fgWhite}
				}

				if slashingPct > slashingPct50 {
					return text.Colors{bgYellow, fgBlack}
				}

				return valuesRowFgColor
			}

			for _, column := range row {
				switch value := column.(type) {
				case int, int16, int32, int64:
					return valuesRowFgColor
				case bool:
					return valuesRowFgColor
				case string:
					if _, err := strconv.Atoi(value); err == nil {
						return valuesRowFgColor
					}
				}
			}

			return text.Colors{fgBlack, bgWhite}
		}

		return handler
	}()

	tb.writer.SetRowPainter(painter)
}
