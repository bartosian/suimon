package tablebuilder

import (
	"errors"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/tablebuilder/tables"
	"github.com/bartosian/sui_helpers/suimon/internal/core/gateways/cligw"
)

type Builder struct {
	tableType  enums.TableType
	hosts      []host.Host
	cliGateway *cligw.Gateway
	writer     table.Writer
	config     *tables.TableConfig
}

// NewBuilder creates a new instance of the table builder, using the CLI gateway
func NewBuilder(tableType enums.TableType, hosts []host.Host, cliGateway *cligw.Gateway) *Builder {
	tableWR := table.NewWriter()
	tableWR.SetOutputMirror(os.Stdout)

	return &Builder{
		tableType:  tableType,
		hosts:      hosts,
		cliGateway: cliGateway,
		writer:     tableWR,
	}
}

// setColumns sets the column configurations for the table builder based on the configuration in the builder's table config
func (tb *Builder) setColumns() {
	var columnsConfig []table.ColumnConfig

	for _, column := range tb.config.Columns {
		columnsConfig = append(columnsConfig, *column.Config)
	}

	tb.writer.SetColumnConfigs(columnsConfig)
}

// setRows sets the rows of the table builder based on the configuration in the builder's table config
func (tb *Builder) setRows() error {
	rowsConfig := tb.config.Rows
	columnsConfig := tb.config.Columns
	itemsCount := tb.config.RowsCount
	columnsPerRow := len(rowsConfig[0])

	for itemIndex := 0; itemIndex < itemsCount; itemIndex++ {
		for rowIndex, columns := range rowsConfig {
			header := tables.NewRow(true, false, columnsPerRow, true, text.AlignCenter)
			footer := tables.NewRow(false, false, columnsPerRow, true, text.AlignCenter)
			row := tables.NewRow(false, true, columnsPerRow, true, text.AlignCenter)

			for _, columnName := range columns {
				columnConfig, ok := columnsConfig[columnName]
				if !ok {
					tb.cliGateway.Errorf("column %s not found", columnName)

					return errors.New("column not found")
				}

				columnValue := columnConfig.Values[itemIndex]

				header.AppendValue(columnName.ToString())
				row.AppendValue(columnValue)
				footer.PrependValue(tables.EmptyValue)
			}

			for columnIdx := len(columns); columnIdx < columnsPerRow; columnIdx++ {
				header.PrependValue(tables.EmptyValue)
				footer.PrependValue(tables.EmptyValue)
				row.PrependValue(tables.EmptyValue)
			}

			if itemIndex == 0 && rowIndex == 0 {
				tb.writer.AppendHeader(header.Values, header.Config)
				tb.writer.AppendFooter(footer.Values, footer.Config)
			} else if len(rowsConfig) > 1 && (rowIndex%2 == 1 || itemIndex > 0 && rowIndex%2 == 0) {
				tb.writer.AppendRow(header.Values, header.Config)
			}

			tb.writer.AppendRow(row.Values, row.Config)
			tb.writer.AppendSeparator()
		}
	}

	return nil
}

// setStyle sets the style for the table builder based on the configuration in the builder's table config
func (tb *Builder) setStyle() {
	tb.writer.SetTitle(tb.config.Name)
	tb.writer.SetStyle(tb.config.Style)

	tb.setColors()
}

// setColors sets the row colors for the table builder based on the current state of the table
func (tb *Builder) setColors() {
	var (
		fgWhite  = text.FgWhite
		bgWhite  = text.BgWhite
		bgHiBlue = text.BgHiBlue
		fgBlack  = text.FgBlack
		bgRed    = text.BgRed
		bgYellow = text.BgYellow
	)

	var painter = func() func(row table.Row) text.Colors {
		valuesRowFgColor := text.Colors{fgWhite}
		bgColor := []text.Color{bgWhite, bgHiBlue, bgHiBlue, bgWhite}
		currentColor := 0

		var handler = func(row table.Row) text.Colors {
			switch tb.tableType {
			case enums.TableTypeValidatorReports:
				valueString, ok := row[1].(string)
				if !ok {
					return valuesRowFgColor
				}

				slashingPct, err := strconv.ParseFloat(valueString, 64)
				if err != nil {
					return valuesRowFgColor
				}

				if slashingPct > 100 {
					return text.Colors{bgRed, fgWhite}
				}

				if slashingPct > 50 {
					return text.Colors{bgYellow, fgBlack}
				}

				return valuesRowFgColor
			default:
				for _, column := range row {
					switch value := column.(type) {
					case int:
						return valuesRowFgColor
					case string:
						if _, err := strconv.Atoi(value); err == nil {
							return valuesRowFgColor
						}
					}
				}

				colors := text.Colors{fgBlack, bgColor[currentColor]}

				currentColor++
				if currentColor > 3 {
					currentColor = 0
				}

				return colors
			}
		}

		return handler
	}()

	tb.writer.SetRowPainter(painter)
}
