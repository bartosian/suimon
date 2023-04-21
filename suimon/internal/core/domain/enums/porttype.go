package enums

type PortType int

const (
	PortTypeRPC PortType = iota
	PortTypeMetrics
	PortTypePeer
)
