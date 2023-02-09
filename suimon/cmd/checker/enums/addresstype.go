package enums

//go:generate go run github.com/dmarkham/enumer -type=AddressType -json -transform=snake-upper -output=./addresstype.gen.go
type AddressType int

const (
	AddressTypeIP AddressType = iota
	AddressTypeDomain
)
