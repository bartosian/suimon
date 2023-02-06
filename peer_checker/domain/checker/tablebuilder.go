package checker

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

const noDataReceived = "🔴 no data"

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
	tb.builder.AppendHeader(table.Row{"#", "Peer", "Port", "Country", "Total TXS"}, rowConfigAutoMerge)

	for idx, peer := range data {
		var totalTransactionsNumber any = noDataReceived

		if peer.TotalTransactionNumber != nil {
			totalTransactionsNumber = *peer.TotalTransactionNumber
		}

		tb.builder.AppendRow(table.Row{idx + 1, peer.Address, peer.Port, peer.Location.String(), totalTransactionsNumber}, rowConfigAutoMerge)
		tb.builder.AppendSeparator()
	}

	tb.builder.Render()
}
