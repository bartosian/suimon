package checker

func (checker *Checker) DrawTable() {
	checker.tableBuilder.BuildTable(checker.Peers)
}
