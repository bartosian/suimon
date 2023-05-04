package tablebuilder

// Render sets the rows, columns, and style for the Builder and then renders the table.
func (tb *Builder) Render() error {
	if err := tb.setRows(); err != nil {
		return err
	}

	tb.setColumns()
	tb.setStyle()
	tb.writer.Render()

	return nil
}
