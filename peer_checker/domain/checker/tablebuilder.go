package checker

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type TableBuilder struct {
	builder table.Writer
}

func NewTableBuilder() *TableBuilder {
	tableWR := table.NewWriter()

	tableWR.SetOutputMirror(os.Stdout)
	tableWR.SetStyle(table.StyleLight)

	return &TableBuilder{
		builder: tableWR,
	}
}

func (tb *TableBuilder) BuildTable(data []Peer) {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	tb.builder.AppendHeader(table.Row{"#", "Peer", "Port", "Country"}, rowConfigAutoMerge)

	for idx, peer := range data {
		tb.builder.AppendRow(table.Row{idx + 1, peer.Address, peer.Port, peer.Location.String()}, rowConfigAutoMerge)
		tb.builder.AppendSeparator()
	}

	tb.builder.Render()
}
