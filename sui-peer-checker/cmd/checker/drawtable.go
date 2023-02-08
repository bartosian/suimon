package checker

func (checker *Checker) DrawTable() {
	checker.tableBuilderRPC.Build()
	checker.tableBuilderPeer.Build()
}
