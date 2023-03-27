package enums

type PortType int

const (
	PortTypeUndefined PortType = iota
	PortTypeRPC
	PortTypeMetrics
	PortTypePeer
)
