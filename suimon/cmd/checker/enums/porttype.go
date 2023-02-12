package enums

//go:generate go run github.com/dmarkham/enumer -type=PortType -json -transform=snake-upper -output=./porttype.gen.go
type PortType int

const (
	PortTypeUndefined PortType = iota
	PortTypeRPC
	PortTypeMetrics
	PortTypePeer
)
