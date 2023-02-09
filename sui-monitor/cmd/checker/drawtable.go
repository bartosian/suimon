package checker

func (checker *Checker) DrawTable() {
	checker.tableBuilderRPC.Build()
	checker.tableBuilderNode.Build()
	checker.tableBuilderPeer.Build()
}
