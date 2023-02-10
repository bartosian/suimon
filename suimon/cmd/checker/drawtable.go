package checker

func (checker *Checker) DrawTable() {
	if checker.suimonConfig.MonitorsConfig.RPCTable.Display {
		checker.tableBuilderRPC.Build()
	}
	if checker.suimonConfig.MonitorsConfig.NodeTable.Display {
		checker.tableBuilderNode.Build()
	}
	if checker.suimonConfig.MonitorsConfig.PeersTable.Display {
		checker.tableBuilderPeer.Build()
	}
}
