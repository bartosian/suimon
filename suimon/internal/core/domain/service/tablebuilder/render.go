package tablebuilder

// Render sets the rows, columns, and style for the Builder and then renders the table.
func (tb *Builder) Render() {
	tb.setRows()
	tb.setColumns()
	tb.setStyle()

	tb.writer.Render()
}
